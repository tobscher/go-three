package three

import (
	"errors"
	"fmt"
	"log"

	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glh "github.com/tobscher/glh"
)

type Renderer struct {
	Width       int
	Height      int
	vertexArray gl.VertexArray
	window      *glfw.Window
}

func NewRenderer(width, height int, title string) (*Renderer, error) {
	// Error callback
	glfw.SetErrorCallback(errorCallback)

	// Init glfw
	if !glfw.Init() {
		return nil, errors.New("Could not initialise GLFW.")
	}

	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	// Create window
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, err
	}
	window.SetKeyCallback(keyCallback)
	window.MakeContextCurrent()

	// Use vsync
	glfw.SwapInterval(1)

	// Init glew
	if gl.Init() != 0 {
		return nil, errors.New("Could not initialise glew.")
	}
	gl.GetError()

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	// Vertex buffers
	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()

	renderer := Renderer{vertexArray: vertexArray, window: window, Width: width, Height: height}
	return &renderer, nil
}

func (r *Renderer) SetSize(width, height int) {
	r.Width = width
	r.Height = height
}

// Generate buffers on the fly if they aren't generated already
// Check if underlying data needs updating.
func (r *Renderer) Render(scene *Scene, camera *PerspectiveCamera) {
	width, height := r.window.GetFramebufferSize()
	gl.Viewport(0, 0, width, height)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, element := range scene.objects {
		program := element.material.Program()
		if program == nil {
			program = createProgram(element)
			element.material.SetProgram(program)
		}
		program.use()

		// Is already inverted by multiplier
		view := camera.Transform.modelMatrix()
		projection := camera.projectionMatrix
		MVP := projection.Mul4(view).Mul4(element.Transform.modelMatrix())

		// Set model view projection matrix
		program.uniforms["MVP"].apply(MVP)

		if c, ok := element.material.(Colored); ok {
			if c.Color() != nil {
				program.uniforms["diffuse"].apply(c.Color())
			}
		}

		var toDisable []gl.AttribLocation

		for _, attribute := range program.attributes {
			toDisable = append(toDisable, attribute.enableFor(element))
		}

		t, ok := element.material.(Wireframed)
		if ok {
			if t.Wireframe() {
				gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
			} else {
				gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
			}
		}

		gl.DrawArrays(gl.TRIANGLES, 0, len(element.geometry.Vertices()))

		for _, location := range toDisable {
			location.DisableArray()
		}
	}

	r.window.SwapBuffers()
	glfw.PollEvents()
}

func (r *Renderer) ShouldClose() bool {
	return r.window.ShouldClose()
}

func createProgram(mesh *Mesh) *Program {
	program := NewProgram()
	material := mesh.material
	geometry := mesh.geometry

	// Attributes
	program.attributes["vertex"] = NewAttribute(0, 3, newVertexBuffer(geometry))

	var feature ProgramFeature
	if c, cOk := material.(Colored); cOk {
		if c.Color() != nil {
			feature = COLOR
		}
	}

	// Let geometry return UVs
	if t, tOk := material.(Textured); tOk {
		if t.Texture() != nil {
			program.attributes["texture"] = NewAttribute(1, 2, newUvBuffer(len(geometry.Vertices()), t))
			feature = TEXTURE
		}
	}

	program.Load(MakeProgram(feature))

	// Uniforms
	program.uniforms["MVP"] = NewUniform(program, "MVP")
	program.uniforms["diffuse"] = NewUniform(program, "diffuse")

	return program
}

func newVertexBuffer(geometry Shape) *buffer {
	result := []float32{}

	for _, vertex := range geometry.Vertices() {
		result = append(result, vertex.X(), vertex.Y(), vertex.Z())
	}

	b := NewBuffer(result)
	return &b
}

// Do not use color value per vertex
// instead use uniform for diffuse color
func newColorBuffer(count int, material Colored) *buffer {
	result := []float32{}

	for i := 0; i < count; i++ {
		color := material.Color()
		result = append(result, color.R(), color.G(), color.B())
	}

	b := NewBuffer(result)
	return &b
}

func newUvBuffer(count int, material Textured) *buffer {
	result := []float32{}

	for i := 0; i < 6; i++ {
		result = append(result,
			1, 1,
			0, 0,
			1, 0,

			1, 1,
			0, 1,
			0, 0,
		)
	}

	// Invert V because we're using a compressed texture
	for i := 1; i < len(result); i += 2 {
		result[i] = 1.0 - result[i]
	}

	b := NewBuffer(result)
	return &b
}

func (r *Renderer) Unload(s *Scene) {
	log.Println("Cleaning up...")

	for _, element := range s.objects {
		program := element.material.Program()
		program.unload()
	}

	r.vertexArray.Delete()
	glfw.Terminate()
}

func (r *Renderer) OpenGLSentinel() {
	glh.OpenGLSentinel()
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}
