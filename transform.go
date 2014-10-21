package three

import (
	"github.com/go-gl/mathgl/mgl32"
)

type transform struct {
	position   mgl32.Vec3
	rotation   mgl32.Vec3
	quaternion mgl32.Quat
	scale      mgl32.Vec3

	multiplier float32

	up      mgl32.Vec3
	right   mgl32.Vec3
	forward mgl32.Vec3

	matrix mgl32.Mat4
}

func NewTransform(multiplier float32) transform {
	return transform{
		position:   mgl32.Vec3{0, 0, 0},
		rotation:   mgl32.Vec3{0, 0, 0},
		quaternion: mgl32.QuatIdent(),
		scale:      mgl32.Vec3{1, 1, 1},

		multiplier: multiplier,

		up:      mgl32.Vec3{0, 1, 0},
		right:   mgl32.Vec3{1, 0, 0},
		forward: mgl32.Vec3{0, 0, -1},

		matrix: mgl32.Ident4(),
	}
}

func (t *transform) SetPosition(x, y, z float32) {
	t.position = mgl32.Vec3{x, y, z}

	t.matrix[12] = x * t.multiplier
	t.matrix[13] = y * t.multiplier
	t.matrix[14] = z * t.multiplier
}

func (t *transform) Scale(x, y, z float32) {
	t.scale = mgl32.Vec3{x, y, z}

	t.matrix[0] = x * t.multiplier
	t.matrix[5] = y * t.multiplier
	t.matrix[10] = z * t.multiplier
}

func (t *transform) RotateX(angle float32) {
	v1 := mgl32.Vec3{1, 0, 0}
	t.rotateOnAxis(v1, angle)
}

func (t *transform) RotateY(angle float32) {
	v1 := mgl32.Vec3{0, 1, 0}
	t.rotateOnAxis(v1, angle)
}

func (t *transform) RotateZ(angle float32) {
	v1 := mgl32.Vec3{0, 0, 1}
	t.rotateOnAxis(v1, angle)
}

func (t *transform) rotateOnAxis(axis mgl32.Vec3, angle float32) {
	q1 := mgl32.QuatRotate(angle*t.multiplier, axis)
	t.matrix = t.matrix.Mul4(q1.Mat4())
}

func (t *transform) LookAt(x, y, z float32) {
	target := mgl32.Vec3{x, y, z}

	t.matrix = mgl32.LookAtV(
		t.position,
		target,
		t.up,
	)
}

func (t *transform) ModelMatrix() mgl32.Mat4 {
	return t.matrix
}
