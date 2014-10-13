package three

import "github.com/go-gl/mathgl/mgl32"

type persepectiveCamera struct {
	Position         mgl32.Vec3
	Up               mgl32.Vec3
	projectionMatrix mgl32.Mat4
	viewMatrix       mgl32.Mat4
}

func NewPerspectiveCamera(fov, aspect, near, far float32) persepectiveCamera {
	matrix := mgl32.Perspective(fov, aspect, near, far)
	defaultPosition := mgl32.Vec3{0.0, 0.0, 0.0}
	defaultUp := mgl32.Vec3{0.0, 1.0, 0.0}
	camera := persepectiveCamera{projectionMatrix: matrix, Position: defaultPosition, Up: defaultUp}
	return camera
}

func (pc *persepectiveCamera) LookAt(target mgl32.Vec3) {
	pc.viewMatrix = mgl32.LookAt(
		pc.Position.X(), pc.Position.Y(), pc.Position.Z(),
		target.X(), target.Y(), target.Z(),
		pc.Up.X(), pc.Up.Y(), pc.Up.Z())
}
