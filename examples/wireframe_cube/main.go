package main

import (
	three "github.com/tobscher/go-three"
	"log"
)

const (
	fov    = 45.0
	width  = 640
	height = 480
	near   = 0.1
	far    = 100
)

func main() {
	renderer, err := three.NewRenderer(width, height, "Wireframe Cube")
	if err != nil {
		log.Fatal(err)
	}

	scene := three.NewScene()
	camera := three.NewPerspectiveCamera(fov, width/height, near, far)
	camera.SetPosition(4.0, 3.0, 4.0)
	camera.LookAt(0, 0, 0)

	box := three.NewCubeGeometry(1)
	blue := three.NewMeshBasicMaterial()
	blue.SetColor(&three.Color{0.0, 0.0, 1.0})
	blue.SetWireframe(true)

	mesh := three.NewMesh(box, blue)

	scene.Add(&mesh)

	for !renderer.ShouldClose() {
		renderer.Render(scene, camera)
	}

	renderer.Unload(&scene)

	renderer.OpenGLSentinel()
}
