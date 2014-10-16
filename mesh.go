package three

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	geometry          Geometry
	material          Material
	position          mgl32.Vec3
	scale             mgl32.Vec3
	translationMatrix mgl32.Mat4
	rotationMatrix    mgl32.Mat4
	scaleMatrix       mgl32.Mat4
}

func NewMesh(geometry Geometry, material Material) Mesh {
	return Mesh{
		geometry:          geometry,
		material:          material,
		position:          mgl32.Vec3{0, 0, 0},
		scale:             mgl32.Vec3{1, 1, 1},
		translationMatrix: mgl32.Ident4(),
		rotationMatrix:    mgl32.Ident4(),
		scaleMatrix:       mgl32.Ident4(),
	}
}

func (m *Mesh) SetPosition(x, y, z float32) {
	m.position = mgl32.Vec3{x, y, z}

	m.translationMatrix[12] = x
	m.translationMatrix[13] = y
	m.translationMatrix[14] = z
}

func (m *Mesh) Scale(x, y, z float32) {
	m.scale = mgl32.Vec3{x, y, z}

	m.scaleMatrix[0] = x
	m.scaleMatrix[5] = y
	m.scaleMatrix[10] = z
}

func (m *Mesh) ModelMatrix() mgl32.Mat4 {
	return m.translationMatrix.Mul4(m.rotationMatrix).Mul4(m.scaleMatrix)
}
