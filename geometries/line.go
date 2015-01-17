package geometries

import (
	"github.com/go-gl/mathgl/mgl32"
	three "github.com/tobscher/go-three"
)

type Line struct {
	three.Geometry

	from mgl32.Vec3
	to   mgl32.Vec3
}

func NewLine(from, to mgl32.Vec3) *Line {
	line := Line{
		from: from,
		to:   to,
	}

	vertices := []mgl32.Vec3{
		from,
		to,
	}

	line.SetVertices(vertices)

	return &line
}
