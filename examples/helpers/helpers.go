package helpers

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tobscher/go-three"
	"github.com/tobscher/go-three/geometries"
)

func AddAxes(scene *three.Scene) {
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
}
