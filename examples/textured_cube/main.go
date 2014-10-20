package main

import (
	"log"

	three "github.com/tobscher/go-three"
)

const (
	fov    = 75.0
	width  = 640
	height = 480
	near   = 1
	far    = 10000
)

func main() {
	renderer, err := three.NewRenderer(width, height, "Application Name")
	if err != nil {
		log.Fatal(err)
	}

	scene := three.NewScene()
	camera := three.NewPerspectiveCamera(fov, width/height, near, far)
	camera.SetPosition(0, 0, 1000)

	box := three.NewCubeGeometry(200)
	texture := three.NewMeshBasicMaterial()
	t, err := three.NewTexture("textures/uvgrid01.dds")
	if err != nil {
		log.Panic(err)
	}
	texture.SetTexture(t)

	mesh := three.NewMesh(box, texture)

	scene.Add(&mesh)

	for !renderer.ShouldClose() {
		mesh.Transform.RotateX(0.01)
		mesh.Transform.RotateY(0.02)
		renderer.Render(scene, camera)
	}

	renderer.Unload(&scene)

	renderer.OpenGLSentinel()
}
