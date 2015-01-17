package main

import (
	"log"
	"runtime"

	three "github.com/tobscher/go-three"
	"github.com/tobscher/go-three/loaders"
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
		Title:      "Example - Obj Loading",
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
	camera.Transform.SetPosition(4.0, 3.0, 4.0)
	camera.Transform.LookAt(0, 0, 0)

	box, err := loaders.LoadFromObj("obj/suzanne.obj")
	if err != nil {
		log.Fatal(err)
	}

	grey := three.NewBasicMaterial()
	grey.SetColor(&three.Color{0.5, 0.5, 0.5})

	mesh := three.NewMesh(box, grey)

	scene.Add(mesh)

	transform := mesh.Transform()
	for !window.ShouldClose() {
		transform.RotateX(0.01)
		transform.RotateY(0.02)

		renderer.Render(scene, camera)
	}

	renderer.Unload(scene)
	renderer.OpenGLSentinel()
}
