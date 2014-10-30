package geometries

import (
	"github.com/go-gl/mathgl/mgl32"
	three "github.com/tobscher/go-three"
)

type Tetrahedron struct {
	geometry three.Geometry
}

func NewTetrahedron() *Tetrahedron {
	tetrahedron := Tetrahedron{}

	vertices := []mgl32.Vec3{
		mgl32.Vec3{1.0, 1.0, 1.0},   // index 0
		mgl32.Vec3{-1.0, -1.0, 1.0}, // index 1
		mgl32.Vec3{-1.0, 1.0, -1.0}, // index 2
		mgl32.Vec3{1.0, -1.0, -1.0}, // index 3
	}

	faces := []*three.Face{
		three.NewFace(0, 1, 2),
		three.NewFace(3, 0, 1),
	}

	tetrahedron.geometry.SetVertices(vertices)
	tetrahedron.geometry.SetFaces(faces)

	return &tetrahedron
}

// Vertices returns the list of used vertices to create a tetrahedron geometry.
func (t *Tetrahedron) Vertices() []mgl32.Vec3 {
	return t.geometry.Vertices()
}

// Faces returns the list of used faces to display a tetrahedron geometry.
func (t *Tetrahedron) Faces() []*three.Face {
	return t.geometry.Faces()
}

// VertexUvs returns the uv mapping for each vertex.
func (t *Tetrahedron) UVs() []mgl32.Vec2 {
	return t.geometry.UVs()
}
