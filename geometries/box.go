package geometries

import (
	"github.com/go-gl/mathgl/mgl32"
	three "github.com/tobscher/go-three"
)

// Box defines a box geometry consisting of 6 faces
type Box struct {
	geometry three.Geometry

	width  float32
	height float32
	depth  float32
}

// NewBox creates a new Box with the given width, height and depth.
// This method will generate the required vertices and its uv mappings.
func NewBox(width, height, depth float32) *Box {
	box := Box{
		width:  width,
		height: height,
		depth:  depth,
	}

	var vertices []mgl32.Vec3
	uvs := boxUvs()

	halfWidth := width / 2.0
	halfHeight := height / 2.0
	halfDepth := depth / 2.0

	// Bottom plane
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 - halfDepth},
	)...)

	// Side 1
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 - halfDepth},
	)...)

	// Side 2
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 + halfDepth},
	)...)

	// // Side 3
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 - halfDepth},
	)...)

	// // Side 4
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 + halfDepth},
	)...)

	// Top plane
	vertices = append(vertices, buildPlane(
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 + halfDepth},
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 - halfDepth},
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 - halfDepth},
	)...)

	box.geometry.SetVertices(vertices)
	box.geometry.SetUVs(uvs)

	return &box
}

// NewCube generates a new Box for the given side.
// Vertices and VertexUvs will be created accordingly.
func NewCube(size float32) *Box {
	return NewBox(size, size, size)
}

// Vertices returns the list of used vertices to create a box geometry.
func (b *Box) Vertices() []mgl32.Vec3 {
	return b.geometry.Vertices()
}

// VertexUvs returns the uv mapping for each vertex.
func (b *Box) UVs() []mgl32.Vec2 {
	return b.geometry.UVs()
}

func buildPlane(v1, v2, v3, v4 mgl32.Vec3) []mgl32.Vec3 {
	return []mgl32.Vec3{
		v1,
		v4,
		v3,
		v1,
		v2,
		v4,
	}
}

func boxUvs() []mgl32.Vec2 {
	result := []mgl32.Vec2{}

	for i := 0; i < 6; i++ {
		result = append(result,
			mgl32.Vec2{1, 1},
			mgl32.Vec2{0, 0},
			mgl32.Vec2{1, 0},

			mgl32.Vec2{1, 1},
			mgl32.Vec2{0, 1},
			mgl32.Vec2{0, 0},
		)
	}

	return result
}
