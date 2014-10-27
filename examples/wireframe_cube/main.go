package main

import (
	three "github.com/tobscher/go-three"
	"github.com/tobscher/go-three/geometries"
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
	camera.Transform.SetPosition(4.0, 3.0, 4.0)
	camera.Transform.LookAt(0, 0, 0)

	box := geometries.NewCube(1)
	blue := three.NewBasicMaterial()
	blue.SetColor(&three.Color{0.0, 0.0, 1.0})
	blue.SetWireframe(true)

	mesh := three.NewMesh(box, blue)

	scene.Add(&mesh)

	blue.SetColor(&three.Color{1.0, 0.0, 0.0})

	for !renderer.ShouldClose() {
		mesh.Transform.RotateX(0.01)
		mesh.Transform.RotateY(0.02)

		renderer.Render(scene, camera)
	}

	renderer.Unload(scene)

	renderer.OpenGLSentinel()
}
