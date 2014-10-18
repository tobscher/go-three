package main

import (
	"github.com/go-gl/mathgl/mgl32"
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
	camera.Position = mgl32.Vec3{4.0, 3.0, 4.0}
	camera.LookAt(mgl32.Vec3{0.0, 0.0, 0.0})

	box := three.NewCubeGeometry(1)
	blue := three.NewMeshBasicMaterial().
		SetColor(&three.Color{0.0, 0.0, 1.0}).
		SetWireframe(true)

	mesh := three.NewMesh(box, blue)

	scene.Add(&mesh)

	for !renderer.ShouldClose() {
		renderer.Render(scene, camera)
	}

	renderer.Unload(&scene)

	renderer.OpenGLSentinel()
}
