package three

import "github.com/go-gl/mathgl/mgl32"

type transform struct {
	position   mgl32.Vec3
	rotation   mgl32.Vec3
	quaternion mgl32.Quat
	scale      mgl32.Vec3

	translationMatrix mgl32.Mat4
	rotationMatrix    mgl32.Mat4
	scaleMatrix       mgl32.Mat4
}

func NewTransform() transform {
	return transform{
		position:   mgl32.Vec3{0, 0, 0},
		rotation:   mgl32.Vec3{0, 0, 0},
		quaternion: mgl32.QuatIdent(),
		scale:      mgl32.Vec3{1, 1, 1},

		translationMatrix: mgl32.Ident4(),
		rotationMatrix:    mgl32.Ident4(),
		scaleMatrix:       mgl32.Ident4(),
	}
}

func (t *transform) SetPosition(x, y, z float32) {
	t.position = mgl32.Vec3{x, y, z}

	t.translationMatrix[12] = x
	t.translationMatrix[13] = y
	t.translationMatrix[14] = z
}

func (t *transform) Scale(x, y, z float32) {
	t.scale = mgl32.Vec3{x, y, z}

	t.scaleMatrix[0] = x
	t.scaleMatrix[5] = y
	t.scaleMatrix[10] = z
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
	t.quaternion = mgl32.QuatRotate(angle, axis)
	t.rotationMatrix = t.rotationMatrix.Mul4(t.quaternion.Mat4())
}

func (t *transform) ModelMatrix() mgl32.Mat4 {
	return t.translationMatrix.Mul4(t.rotationMatrix).Mul4(t.scaleMatrix)
}
