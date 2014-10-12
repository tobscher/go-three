package main

import three "github.com/tobscher/go-three"

const (
	fov    = 75
	width  = 640
	height = 480
	near   = 1
	far    = 10000
)

func main() {
	scene := three.NewScene()
	camera := three.NewPerspectiveCamera(fov, width/height, near, far)
	camera.Position.SetZ(1000)

	geometry := three.NewBoxGeometry(200, 200, 200)
	material := three.NewMeshBasicMaterial(0xff0000)

	mesh := three.NewMesh(geometry, material)
	scene.Add(mesh)

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
