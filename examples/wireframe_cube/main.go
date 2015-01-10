package main

import (
	"log"
	"runtime"

	three "github.com/tobscher/go-three"
	"github.com/tobscher/go-three/geometries"
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
		Title:      "Example - Wireframe Cube",
		Fullscreen: false,
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

	box := geometries.NewCube(1)
	red := three.NewBasicMaterial()
	red.SetColor(&three.Color{1.0, 0.0, 0.0})
	red.SetWireframe(true)

	mesh := three.NewMesh(box, red)

	scene.Add(&mesh)

	for !window.ShouldClose() {
		mesh.Transform.RotateX(0.01)
		mesh.Transform.RotateY(0.02)

		renderer.Render(scene, camera)
	}

	renderer.Unload(scene)

	renderer.OpenGLSentinel()
}
