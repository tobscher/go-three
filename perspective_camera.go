package three

import (
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

// PerspectiveCamera has information about the transformation of the camera
// and it's projection matrix.
//
// Note: The underlying matrix for the transform structure must be inverted.
type PerspectiveCamera struct {
	Transform        *Transform
	projectionMatrix mgl32.Mat4
}

// CameraSettings holds information to construct a new camera object.
type CameraSettings struct {
	FOV  float32
	Near float32
	Far  float32
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
	logger.Debug("Creating new perspective camera")
	logger.Trace(fmt.Sprintf("-- FOV: %v", fov))
	logger.Trace(fmt.Sprintf("-- Aspect: %v", aspect))
	logger.Trace(fmt.Sprintf("-- Near: %v", near))
	logger.Trace(fmt.Sprintf("-- Far: %v", far))

	matrix := makePerspective(fov, aspect, near, far)

	camera := PerspectiveCamera{
		projectionMatrix: matrix,
		Transform:        NewTransform(),
	}

	return &camera
}

func makePerspective(fov, aspect, near, far float32) mgl32.Mat4 {
	ymax := near * float32(math.Tan(float64(mgl32.DegToRad(fov*0.5))))
	ymin := -ymax
	xmin := ymin * aspect
	xmax := ymax * aspect

	return mgl32.Frustum(xmin, xmax, ymin, ymax, near, far)
}
