package three

import "github.com/go-gl/mathgl/mgl32"

// Geometry is the interface which defines the shape of a
// 3D object.
type Shape interface {
	Vertices() []mgl32.Vec3
}

type Geometry struct {
	Vertices []mgl32.Vec3
}
