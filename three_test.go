package three

import (
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"log"
)

func RunIn3DContext(example func()) {
	if !glfw.Init() {
		log.Panic("glfw Error")
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	w, h := 100, 100
	window, err := glfw.CreateWindow(w, h, "Test", nil, nil)
	if err != nil {
		log.Panic(err)
	}
	window.MakeContextCurrent()

	if gl.Init() != 0 {
		log.Panic("gl error")
	}
	gl.GetError()

	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()

	example()
}
