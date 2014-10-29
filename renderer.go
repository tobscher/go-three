package three

import (
	"errors"
	"fmt"
	"log"
	"runtime"

	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glh "github.com/tobscher/glh"
)

// Renderer handles mesh rendering to the window.
type Renderer struct {
	Width       int
	Height      int
	vertexArray gl.VertexArray
	window      *glfw.Window
}

// NewRenderer creates a new Renderer with the given window size and title.
func NewRenderer(width, height int, title string) (*Renderer, error) {
	runtime.LockOSThread()

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

// Render renders the given scene with the given camera to the window.
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
			toDisable = append(toDisable, attribute.enableFor())
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

// ShouldClose indicates if the window should be closed.
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
			program.attributes["texture"] = NewAttribute(1, 2, newUvBuffer(geometry))
			feature = TEXTURE
		}
	}

	program.Load(MakeProgram(feature))

	// Uniforms
	program.uniforms["MVP"] = NewUniform(program, "MVP")
	program.uniforms["diffuse"] = NewUniform(program, "diffuse")

	return program
}

func newVertexBuffer(geometry Shape) *Buffer {
	result := []float32{}

	for _, vertex := range geometry.Vertices() {
		result = append(result, vertex.X(), vertex.Y(), vertex.Z())
	}

	b := NewBuffer(result)
	return &b
}

func newUvBuffer(geometry Shape) *Buffer {
	result := []float32{}

	for _, uv := range geometry.VertexUvs() {
		result = append(result, uv.X(), uv.Y())
	}

	// Invert V because we're using a compressed texture
	for i := 1; i < len(result); i += 2 {
		result[i] = 1.0 - result[i]
	}

	b := NewBuffer(result)
	return &b
}

// Unload deallocates the given scene and all its shader programs.
func (r *Renderer) Unload(s *Scene) {
	log.Println("Cleaning up...")

	for _, element := range s.objects {
		program := element.material.Program()
		program.unload()
	}

	r.vertexArray.Delete()
	glfw.Terminate()
}

// OpenGLSentinel reports any OpenGL related errors.
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
