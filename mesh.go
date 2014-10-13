package three

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	geometry Geometry
	material meshBasicMaterial
	position mgl32.Vec3
}

func NewMesh(geometry Geometry, material meshBasicMaterial) Mesh {
	return Mesh{geometry: geometry, material: material}
}

func (m *Mesh) SetPosition(x, y, z float32) {
	m.position = mgl32.Vec3{x, y, z}
	m.geometry.updateBuffer(m.position)
}
