package three

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Use struct composition
type boxGeometry struct {
	width  float32
	height float32
	depth  float32
}

func NewBoxGeometry(width, height, depth float32) *boxGeometry {
	return &boxGeometry{
		width:  width,
		height: height,
		depth:  depth}
}

func NewCubeGeometry(size float32) *boxGeometry {
	return NewBoxGeometry(size, size, size)
}

func (bg *boxGeometry) generateVertexBuffer() []float32 {
	bufferData := make([]float32, 0)

	halfWidth := bg.width / 2.0
	halfHeight := bg.height / 2.0
	halfDepth := bg.depth / 2.0

	// Bottom plane
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 + halfDepth},
	)...)

	// Side 1
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 - halfDepth},
	)...)

	// Side 2
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 + halfDepth},
	)...)

	// // Side 3
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 - halfDepth},
	)...)

	// // Side 4
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 + halfDepth},
	)...)

	// Top plane
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 + halfDepth},
	)...)

	return bufferData
}

func buildPlane(v1, v2, v3, v4 mgl32.Vec3) []float32 {
	return []float32{
		v1.X(), v1.Y(), v1.Z(),
		v4.X(), v4.Y(), v4.Z(),
		v3.X(), v3.Y(), v3.Z(),
		v1.X(), v1.Y(), v1.Z(),
		v2.X(), v2.Y(), v2.Z(),
		v4.X(), v4.Y(), v4.Z(),
	}
}
