package three

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Transform stores information about the position, rotation and scale
// of an 3D object.
type Transform struct {
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

// NewTransform creates a new Transform struct with defaults.
// The given multiplier can be used to invert the matrix, e.g. camera matrix
// This value should be 1 or -1 (inverted).
//
// Position: 0,0,0
// Rotation: 0,0,0
// Scale:    1,1,1
//
// Up:       0,1,0
// Right:    1,0,0
// Forward:  0,0,-1
func NewTransform(multiplier float32) Transform {
	return Transform{
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

// SetPosition sets the position of the 3D object
// and updates it's matrix accordingly.
func (t *Transform) SetPosition(x, y, z float32) {
	t.position = mgl32.Vec3{x, y, z}

	t.matrix[12] = x * t.multiplier
	t.matrix[13] = y * t.multiplier
	t.matrix[14] = z * t.multiplier
}

// TranslateX moves the object along the x axis by the given units.
// The model matrix is updated accordingly.
func (t *Transform) TranslateX(x float32) {
	t.position[0] += x

	t.matrix[12] = t.position[0]
}

// TranslateY moves the object along the y axis by the given units.
// The model matrix is updated accordingly.
func (t *Transform) TranslateY(y float32) {
	t.position[1] += y

	t.matrix[13] = t.position[1]
}

// TranslateZ moves the object along the z axis by the given units.
// The model matrix is updated accordingly.
func (t *Transform) TranslateZ(z float32) {
	t.position[2] += z

	t.matrix[14] = t.position[2]
}

// Translate moves the object by the given vector.
// The model matrix is updated accordingly.
func (t *Transform) Translate(v mgl32.Vec3) {
	t.position = t.position.Add(v)

	t.matrix[12] = t.position[0]
	t.matrix[13] = t.position[1]
	t.matrix[14] = t.position[2]
}

// Scale sets the scale factor of the 3D object to the given values
// and updates it's matrix accordingly.
func (t *Transform) Scale(x, y, z float32) {
	t.scale = mgl32.Vec3{x, y, z}

	t.matrix[0] = x * t.multiplier
	t.matrix[5] = y * t.multiplier
	t.matrix[10] = z * t.multiplier
}

// RotateX rotates the 3D object by the given angle (in radians) around the x axis.
// The model matrix is updated accordingly.
func (t *Transform) RotateX(angle float32) {
	t.rotation[0] += angle

	v1 := mgl32.Vec3{1, 0, 0}
	t.rotateOnAxis(v1, angle)
}

// RotateY rotates the 3D object by the given angle (in radians) around the x axis.
// The model matrix is updated accordingly.
func (t *Transform) RotateY(angle float32) {
	t.rotation[1] += angle

	v1 := mgl32.Vec3{0, 1, 0}
	t.rotateOnAxis(v1, angle)
}

// RotateZ rotates the 3D object by the given angle (in radians) around the x axis.
// The model matrix is updated accordingly.
func (t *Transform) RotateZ(angle float32) {
	t.rotation[2] += angle

	v1 := mgl32.Vec3{0, 0, 1}
	t.rotateOnAxis(v1, angle)
}

func (t *Transform) rotateOnAxis(axis mgl32.Vec3, angle float32) {
	q1 := mgl32.QuatRotate(angle*t.multiplier, axis)
	t.quaternion = t.quaternion.Mul(q1)

	t.matrix = t.matrix.Mul4(q1.Mat4())
}

// LookAt changes the transformation of the 3D object
// to face the target's position. The model matrix
// will be updated accordingly.
//
// Note: This transformation makes use of the up vector.
func (t *Transform) LookAt(x, y, z float32) {
	target := mgl32.Vec3{x, y, z}

	t.matrix = mgl32.LookAtV(
		t.position,
		target,
		t.up,
	)
}

func (t *Transform) modelMatrix() mgl32.Mat4 {
	return t.matrix
}
