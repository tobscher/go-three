package three

import "github.com/go-gl/mathgl/mgl32"

// PerspectiveCamera has information about the transformation of the camera
// and it's projection matrix.
//
// Note: The underlying matrix for the transform structure must be inverted.
type PerspectiveCamera struct {
	Transform        Transform
	projectionMatrix mgl32.Mat4
}

// NewPerspectiveCamera creates a new perspective camera for the given values.
//
// fov: Field of view in degrees
// aspect: aspect ratio
// near: near clip plane
// far: far clip plane
//
// The cameras transform matrix will be inverted.
func NewPerspectiveCamera(fov, aspect, near, far float32) *PerspectiveCamera {
	matrix := mgl32.Perspective(fov, aspect, near, far)

	camera := PerspectiveCamera{
		projectionMatrix: matrix,
		Transform:        NewTransform(-1),
	}

	return &camera
}
