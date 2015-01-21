package main

import (
	"log"
	"runtime"

	"github.com/go-gl/mathgl/mgl32"
	three "github.com/tobscher/go-three"
)

const (
	fov    = 45.0
	width  = 800
	height = 600
	near   = 0.1
	far    = 100
)

func main() {
	runtime.LockOSThread()

	settings := three.WindowSettings{
		Width:      width,
		Height:     height,
		Title:      "Example - Text",
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

	var size float32 = 15
	var offset float32 = 50.0

	for i := 0; i < 16; i++ {
		white := three.NewBasicMaterial()
		white.SetColor(&three.Color{1.0, 1.0, 1.0})

		regular, err := three.NewFont("../_fonts/Inconsolata-Regular.ttf", int32(15+i))
		if err != nil {
			log.Fatal(err)
		}

		titleGeometry := three.NewTextGeometry("Grumpy wizards", mgl32.Vec2{10, offset}, size, regular)
		title := three.NewText(titleGeometry, white)
		scene.AddText(title)

		offset += size + 5
		size += 2
	}

	camera.Transform.SetPosition(20, 20, 20)
	camera.Transform.LookAt(0, 0, 0)

	for !window.ShouldClose() {
		renderer.Render(scene, camera)
	}

	renderer.Unload(scene)

	renderer.OpenGLSentinel()
}
