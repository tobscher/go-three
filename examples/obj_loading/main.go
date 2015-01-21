package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/mathgl/mgl32"
	three "github.com/tobscher/go-three"
	"github.com/tobscher/go-three/examples/helpers"
	"github.com/tobscher/go-three/loaders"
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
		Title:      "Example - Obj Loading",
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
	camera.Transform.SetPosition(0, 0, 5.0)

	box, err := loaders.LoadFromObj("obj/suzanne.obj")
	if err != nil {
		log.Fatal(err)
	}

	grey := three.NewBasicMaterial()
	grey.SetColor(&three.Color{0.5, 0.5, 0.5})

	helpers.AddAxes(scene)

	mesh := three.NewMesh(box, grey)

	white := three.NewBasicMaterial()
	white.SetColor(&three.Color{1.0, 1.0, 1.0})

	regular, err := three.NewFont("../_fonts/Inconsolata-Regular.ttf", int32(25))
	if err != nil {
		log.Fatal(err)
	}

	fpsGeometry := three.NewTextGeometry(" ", mgl32.Vec2{600, 550}, 25, regular)
	fps := three.NewText(fpsGeometry, white)
	scene.AddText(fps)

	scene.Add(mesh)

	transform := mesh.Transform()
	lastTime := three.GetTime()
	nbFrames := 0

	for !window.ShouldClose() {
		currTime := three.GetTime()
		nbFrames++
		if currTime-lastTime >= 1.0 {
			frames := fmt.Sprintf("%v FPS", nbFrames)
			fps.SetText(frames)

			nbFrames = 0
			lastTime += 1.0
		}

		transform.RotateX(0.01)
		transform.RotateY(0.02)

		renderer.Render(scene, camera)
	}

	renderer.Unload(scene)
	renderer.OpenGLSentinel()
}
