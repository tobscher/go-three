package three

import (
	"errors"
	"fmt"
	"io/ioutil"

	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glh "github.com/tobscher/glh"
)

type Renderer struct {
	Width  int
	Height int
	window *glfw.Window
}

func loadDataFile(filePath string) string {
	content, err := ioutil.ReadFile("shaders/" + filePath)
	if err != nil {
		panic(err)
	}
	return string(content)
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

	renderer := Renderer{window: window, Width: width, Height: height}
	return &renderer, nil
}

func (r *Renderer) SetSize(width, height int) {
	r.Width = width
	r.Height = height
}

func (r *Renderer) Render(scene scene, camera persepectiveCamera) {
	width, height := r.window.GetFramebufferSize()
	gl.Viewport(0, 0, width, height)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, element := range scene.objects {
		element.geometry.Program().Use()

		projection := camera.projectionMatrix
		view := camera.viewMatrix
		MVP := projection.Mul4(view).Mul4(element.ModelMatrix())

		element.geometry.MatrixID().UniformMatrix4fv(false, MVP)

		attribLoc := gl.AttribLocation(0)
		attribLoc.EnableArray()
		element.geometry.Buffer().Bind(gl.ARRAY_BUFFER)
		attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

		colorLoc := gl.AttribLocation(1)
		colorLoc.EnableArray()
		element.material.Buffer(element.geometry.VertexCount()).Bind(gl.ARRAY_BUFFER)
		colorLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

		if element.material.wireframe {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		} else {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
		}

		gl.DrawArrays(gl.TRIANGLES, 0, element.geometry.VertexCount())

		attribLoc.DisableArray()
		colorLoc.DisableArray()
	}
	r.window.SwapBuffers()
	glfw.PollEvents()
}

func (r *Renderer) ShouldClose() bool {
	return r.window.ShouldClose()
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
