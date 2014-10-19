package three

import "github.com/go-gl/mathgl/mgl32"

type persepectiveCamera struct {
	Target           mgl32.Vec3
	Position         mgl32.Vec3
	Up               mgl32.Vec3
	projectionMatrix mgl32.Mat4
	viewMatrix       mgl32.Mat4
}

func NewPerspectiveCamera(fov, aspect, near, far float32) persepectiveCamera {
	matrix := mgl32.Perspective(fov, aspect, near, far)
	viewMatrix := mgl32.Ident4()
	defaultPosition := mgl32.Vec3{0.0, 0.0, 0.0}
	defaultUp := mgl32.Vec3{0.0, 1.0, 0.0}
	defaultTarget := mgl32.Vec3{0, 0, 0}

	camera := persepectiveCamera{
		projectionMatrix: matrix,
		viewMatrix:       viewMatrix,
		Position:         defaultPosition,
		Up:               defaultUp,
		Target:           defaultTarget,
	}
	return camera
}

func (pc *persepectiveCamera) SetUp(x, y, z float32) {
	pc.Up = mgl32.Vec3{x, y, z}
	pc.LookAt(pc.Target.X(), pc.Target.Y(), pc.Target.Z())
}

func (pc *persepectiveCamera) SetPosition(x, y, z float32) {
	pc.Position = mgl32.Vec3{x, y, z}
	pc.LookAt(pc.Target.X(), pc.Target.Y(), pc.Target.Z())
}

func (pc *persepectiveCamera) LookAt(x, y, z float32) {
	pc.Target = mgl32.Vec3{x, y, z}

	pc.viewMatrix = mgl32.LookAt(
		pc.Position.X(), pc.Position.Y(), pc.Position.Z(),
		pc.Target.X(), pc.Target.Y(), pc.Target.Z(),
		pc.Up.X(), pc.Up.Y(), pc.Up.Z())
}
