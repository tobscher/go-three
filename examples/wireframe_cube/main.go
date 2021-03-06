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
		ClearColor: &three.Color{0., 0., 0.4},
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

	scene.Add(mesh)

	var rotX float32 = 0.01
	var rotY float32 = 0.02
	transform := mesh.Transform()
	for !window.ShouldClose() {
		transform.RotateX(rotX)
		transform.RotateY(rotY)

		renderer.Render(scene, camera)
	}

	renderer.Unload(scene)

	renderer.OpenGLSentinel()
}
