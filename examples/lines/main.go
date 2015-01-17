package main

import (
	"log"
	"runtime"

	three "github.com/tobscher/go-three"
	"github.com/tobscher/go-three/examples/helpers"
)

const (
	fov    = 45.0
	width  = 640
	height = 480
	near   = 0.1
	far    = 100
)

func main() {
	runtime.LockOSThread()

	settings := three.WindowSettings{
		Width:      width,
		Height:     height,
		Title:      "Example - Lines",
		Fullscreen: false,
		// ClearColor: &three.Color{0., 0., 0.4},
	}

	window, err := three.NewWindow(settings)
	if err != nil {
		log.Fatal(err)
	}

	renderer, err := three.NewRenderer(window)
	if err != nil {
		log.Fatal(err)
	}

	scene := three.NewScene()
	camera := three.NewPerspectiveCamera(fov, width/height, near, far)

	helpers.AddAxes(scene)

	camera.Transform.SetPosition(20, 20, 20)
	camera.Transform.LookAt(0, 0, 0)

	for !window.ShouldClose() {
		renderer.Render(scene, camera)
	}

	renderer.Unload(scene)

	renderer.OpenGLSentinel()
}
