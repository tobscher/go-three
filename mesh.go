package three

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	vertexBuffer buffer
	colorBuffer  buffer
	uvBuffer     buffer

	geometry Geometry
	material Material

	position   mgl32.Vec3
	rotation   mgl32.Vec3
	quaternion mgl32.Quat
	scale      mgl32.Vec3

	translationMatrix mgl32.Mat4
	rotationMatrix    mgl32.Mat4
	scaleMatrix       mgl32.Mat4
}

func NewMesh(geometry Geometry, material Material) Mesh {
	vertexBuffer := NewBuffer(geometry.generateVertexBuffer())

	m := Mesh{
		vertexBuffer: vertexBuffer,
		geometry:     geometry,
		material:     material,

		position:   mgl32.Vec3{0, 0, 0},
		rotation:   mgl32.Vec3{0, 0, 0},
		quaternion: mgl32.QuatIdent(),
		scale:      mgl32.Vec3{1, 1, 1},

		translationMatrix: mgl32.Ident4(),
		rotationMatrix:    mgl32.Ident4(),
		scaleMatrix:       mgl32.Ident4(),
	}

	if material.Color() != nil {
		colorBuffer := NewBuffer(material.generateColorBuffer(vertexBuffer.vertexCount()))
		m.colorBuffer = colorBuffer
	}

	if material.Texture() != nil {
		uvBuffer := NewBuffer(material.generateUvBuffer(vertexBuffer.vertexCount()))
		m.uvBuffer = uvBuffer
	}

	return m
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

func (m *Mesh) RotateX(angle float32) {
	v1 := mgl32.Vec3{1, 0, 0}
	m.rotateOnAxis(v1, angle)
}

func (m *Mesh) RotateY(angle float32) {
	v1 := mgl32.Vec3{0, 1, 0}
	m.rotateOnAxis(v1, angle)
}

func (m *Mesh) RotateZ(angle float32) {
	v1 := mgl32.Vec3{0, 0, 1}
	m.rotateOnAxis(v1, angle)
}

func (m *Mesh) rotateOnAxis(axis mgl32.Vec3, angle float32) {
	m.quaternion = mgl32.QuatRotate(angle, axis)
	m.rotationMatrix = m.rotationMatrix.Mul4(m.quaternion.Mat4())
}

func (m *Mesh) ModelMatrix() mgl32.Mat4 {
	return m.translationMatrix.Mul4(m.rotationMatrix).Mul4(m.scaleMatrix)
}

func (m *Mesh) ColorBuffer() buffer {
	if m.material.ColorsDirty() {
		m.colorBuffer.update(m.material.generateColorBuffer(m.vertexBuffer.vertexCount()))
		m.material.SetColorsDirty(false)
	}

	return m.colorBuffer
}

func (m *Mesh) UvBuffer() buffer {
	if m.material.TextureDirty() {
		m.uvBuffer.update(m.material.generateUvBuffer(m.vertexBuffer.vertexCount()))
		m.material.SetTextureDirty(false)
	}

	return m.uvBuffer
}
