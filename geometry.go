package three

import "github.com/go-gl/mathgl/mgl32"

// Shape is the interface which defines the shape of a
// 3D object.
type Shape interface {
	Vertices() []mgl32.Vec3
	VertexUvs() []mgl32.Vec2
}

// Geometry is a base struct with fields for Vertices and VertexUvs.
type Geometry struct {
	Vertices  []mgl32.Vec3
	VertexUvs []mgl32.Vec2
}
