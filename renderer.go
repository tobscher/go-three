package three

import (
	"errors"

	gl "github.com/go-gl/gl"
	glh "github.com/tobscher/glh"
)

// Renderer handles mesh rendering to the window.
type Renderer struct {
	window      *Window
	vertexArray gl.VertexArray
}

// NewRenderer creates a new Renderer with the given window size and title.
func NewRenderer(window *Window) (*Renderer, error) {
	// Init glew
	if gl.Init() != 0 {
		return nil, errors.New("Could not initialise glew.")
	}
	gl.GetError()

	if window.Settings.ClearColor != nil {
		color := window.Settings.ClearColor
		gl.ClearColor(
			gl.GLclampf(color.R()),
			gl.GLclampf(color.G()),
			gl.GLclampf(color.B()),
			0.,
		)
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.Enable(gl.CULL_FACE)

	// Vertex buffers
	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()

	renderer := Renderer{
		vertexArray: vertexArray,
		window:      window,
	}
	return &renderer, nil
}

// Render renders the given scene with the given camera to the window.
func (r *Renderer) Render(scene *Scene, camera *PerspectiveCamera) {
	gl.Viewport(0, 0, r.window.Settings.Width, r.window.Settings.Height)
	// Seems to be causing a strange memory leak
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, element := range scene.objects {
		handleElement(camera, element)
	}

	r.window.Swap()
}

func handleElement(camera *PerspectiveCamera, element SceneObject) {
	// Material
	material := element.Material()
	program := material.Program()
	if program == nil {
		program = createProgram(element)
		material.SetProgram(program)
	}
	program.Use()
	defer program.Unuse()

	view := camera.Transform.modelMatrix().Inv()
	projection := camera.projectionMatrix
	MVP := projection.Mul4(view).Mul4(element.Transform().modelMatrix())

	// Set model view projection matrix
	program.uniforms["MVP"].apply(MVP)

	if c, ok := material.(Colored); ok {
		if c.Color() != nil {
			program.uniforms["diffuse"].apply(c.Color())
		}
	}

	if t, ok := material.(Textured); ok {
		texture := t.Texture()
		if texture != nil {
			gl.ActiveTexture(gl.TEXTURE0)
			texture.Bind()
			defer texture.Unbind()
			program.uniforms["texture"].apply(texture)
			program.uniforms["repeat"].apply(texture.Repeat)
		}
	}

	for _, attribute := range program.attributes {
		attribute.enable()
		defer attribute.disable()
		attribute.bindBuffer()
		defer attribute.unbindBuffer()
		attribute.pointer()
		attribute.bindBuffer()
	}

	vertexAttrib := gl.AttribLocation(0)
	vertexAttrib.EnableArray()
	defer vertexAttrib.DisableArray()
	element.VertexBuffer().Bind(gl.ARRAY_BUFFER)
	defer element.VertexBuffer().Unbind(gl.ARRAY_BUFFER)
	vertexAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)

	if t, ok := material.(Wireframed); ok {
		if t.Wireframe() {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		} else {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
		}
	}

	// If index available
	index := element.Index()
	if index != nil && index.count > 0 {
		index.enable()
		defer index.disable()

		gl.DrawElements(gl.TRIANGLES, index.count, gl.UNSIGNED_SHORT, nil)
	} else {
		gl.DrawArrays(element.Mode(), 0, len(element.Geometry().Vertices()))
	}
}

func createProgram(sceneObject SceneObject) *Program {
	program := NewProgram()
	material := sceneObject.Material()

	// Attributes
	var feature ProgramFeature
	if c, cOk := material.(Colored); cOk {
		if c.Color() != nil {
			feature = COLOR
		}
	}

	// Let geometry return UVs
	if t, tOk := material.(Textured); tOk {
		if t.Texture() != nil {
			program.attributes["texture"] = NewAttribute(1, 2, sceneObject.UVBuffer())
			feature = TEXTURE
		}
	}

	program.Load(MakeProgram(feature))

	// Uniforms
	program.uniforms["MVP"] = NewUniform(program, "MVP")
	program.uniforms["diffuse"] = NewUniform(program, "diffuse")

	if t, tOk := material.(Textured); tOk {
		if texture := t.Texture(); texture != nil {
			program.uniforms["texture"] = NewUniform(program, "textureSampler")
			program.uniforms["repeat"] = NewUniform(program, "repeat")
		}
	}

	return program
}

// Unload deallocates the given scene and all its shader programs.
func (r *Renderer) Unload(s *Scene) {
	for _, element := range s.objects {
		program := element.Material().Program()
		program.unload()
	}

	r.vertexArray.Delete()
	r.window.Unload()
}

// OpenGLSentinel reports any OpenGL related errors.
func (r *Renderer) OpenGLSentinel() {
	glh.OpenGLSentinel()
}
