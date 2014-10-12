package main

import (
	"github.com/go-gl/mathgl/mgl32"
	three "github.com/tobscher/go-three"
)

const (
	fov    = 45.0
	width  = 640
	height = 480
	near   = 0.1
	far    = 100
)

func main() {
	scene := three.NewScene()
	camera := three.NewPerspectiveCamera(fov, width/height, near, far)
	camera.Position = mgl32.Vec3{4.0, 3.0, 3.0}
	camera.LookAt(mgl32.Vec3{0.0, 0.0, 0.0})

	geometry := three.NewTriangleGeometry()
	material := three.NewMeshBasicMaterial(0xff0000)

	mesh := three.NewMesh(geometry, material)
	scene.Add(&geometry)

	renderer, err := three.NewRenderer(width, height, "Application Name")
	if err != nil {
		panic(err)
	}

	for !renderer.ShouldClose() {
		mesh.Rotation.X += 0.01
		mesh.Rotation.Y += 0.02

		renderer.Render(scene, camera)
	}

	renderer.OpenGLSentinel()
}