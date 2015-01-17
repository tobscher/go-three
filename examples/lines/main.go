package main

import (
	"log"
	"runtime"

	"github.com/go-gl/mathgl/mgl32"
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

	red := three.NewBasicMaterial()
	red.SetColor(&three.Color{1.0, 0.0, 0.0})

	green := three.NewBasicMaterial()
	green.SetColor(&three.Color{0.0, 1.0, 0.0})

	blue := three.NewBasicMaterial()
	blue.SetColor(&three.Color{0.0, 0.0, 1.0})

	lineX := geometries.NewLine(mgl32.Vec3{0, 0, 0}, mgl32.Vec3{10, 0, 0})
	lineY := geometries.NewLine(mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 10, 0})
	lineZ := geometries.NewLine(mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 0, 10})

	scene.Add(three.NewLine(lineX, red))
	scene.Add(three.NewLine(lineY, green))
	scene.Add(three.NewLine(lineZ, blue))

	camera.Transform.SetPosition(20, 20, 20)
	camera.Transform.LookAt(0, 0, 0)

	for !window.ShouldClose() {
		renderer.Render(scene, camera)
	}

	renderer.Unload(scene)

	renderer.OpenGLSentinel()
}
