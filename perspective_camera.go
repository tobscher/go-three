package three

import "github.com/go-gl/mathgl/mgl32"
import "math"

type perspectiveCamera struct {
	vAngle           float64
	hAngle           float64
	Position         mgl32.Vec3
	projectionMatrix mgl32.Mat4
	viewMatrix       mgl32.Mat4
}

func NewPerspectiveCamera(fov, aspect, near, far float32) perspectiveCamera {
	matrix := mgl32.Perspective(fov, aspect, near, far)
	viewMatrix := mgl32.Ident4()
	defaultPosition := mgl32.Vec3{0.0, 0.0, 0.0}

	camera := perspectiveCamera{
		hAngle:           math.Pi,
		vAngle:           0.0,
		projectionMatrix: matrix,
		viewMatrix:       viewMatrix,
		Position:         defaultPosition,
	}
	return camera
}

func (pc *perspectiveCamera) SetPosition(x, y, z float32) {
	pc.Position = mgl32.Vec3{x, y, z}
	pc.updateViewMatrix()
}

func (pc *perspectiveCamera) LookAt(x, y, z float32) {
	target := mgl32.Vec3{x, y, z}
	pc.viewMatrix = mgl32.LookAtV(
		pc.Position,
		target,
		pc.Up())
}

func (pc *perspectiveCamera) Direction() mgl32.Vec3 {
	dir := mgl32.Vec3{
		float32(math.Cos(pc.vAngle) * math.Sin(pc.hAngle)),
		float32(math.Sin(pc.vAngle)),
		float32(math.Cos(pc.vAngle) * math.Cos(pc.hAngle))}

	return dir
}

func (pc *perspectiveCamera) Up() mgl32.Vec3 {
	right := mgl32.Vec3{
		float32(math.Sin(pc.hAngle - math.Pi/2.0)),
		0.0,
		float32(math.Cos(pc.hAngle - math.Pi/2.0))}

	return right.Cross(pc.Direction())
}

func (pc *perspectiveCamera) updateViewMatrix() {
	pc.viewMatrix = mgl32.LookAtV(pc.Position, pc.Position.Add(pc.Direction()), pc.Up())
}
