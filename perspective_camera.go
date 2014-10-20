package three

import "github.com/go-gl/mathgl/mgl32"

type perspectiveCamera struct {
	Transform        transform
	projectionMatrix mgl32.Mat4
}

func NewPerspectiveCamera(fov, aspect, near, far float32) *perspectiveCamera {
	matrix := mgl32.Perspective(fov, aspect, near, far)

	camera := perspectiveCamera{
		projectionMatrix: matrix,
		Transform:        NewTransform(-1),
	}

	return &camera
}
