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

	uvs := boxUvs()

	halfWidth := width / 2.0
	halfHeight := height / 2.0
	halfDepth := depth / 2.0

	vertices := []mgl32.Vec3{
		// Bottom vertices
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 + halfDepth}, // 1
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 + halfDepth}, // 2
		mgl32.Vec3{0 + halfWidth, 0 - halfHeight, 0 - halfDepth}, // 3
		mgl32.Vec3{0 - halfWidth, 0 - halfHeight, 0 - halfDepth}, // 4

		// Top vertices
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 + halfDepth}, // 5
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 + halfDepth}, // 6
		mgl32.Vec3{0 - halfWidth, 0 + halfHeight, 0 - halfDepth}, // 7
		mgl32.Vec3{0 + halfWidth, 0 + halfHeight, 0 - halfDepth}, // 8
	}

	var faces []*three.Face
	// Bottom
	faces = append(faces, buildFace(0, 1, 2, 3)...)
	// Side 1
	faces = append(faces, buildFace(2, 3, 7, 6)...)
	// Side 2
	faces = append(faces, buildFace(3, 1, 6, 4)...)
	// Side 3
	faces = append(faces, buildFace(0, 2, 5, 7)...)
	// Side 4
	faces = append(faces, buildFace(1, 0, 4, 5)...)
	// Top
	faces = append(faces, buildFace(4, 5, 6, 7)...)

	box.geometry.SetVertices(vertices)
	box.geometry.SetUVs(uvs)
	box.geometry.SetFaces(faces)

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

// Faces returns the list of used faces to display a box geometry.
func (b *Box) Faces() []*three.Face {
	return b.geometry.Faces()
}

// VertexUvs returns the uv mapping for each vertex.
func (b *Box) UVs() []mgl32.Vec2 {
	return b.geometry.UVs()
}

func buildFace(v1, v2, v3, v4 uint16) []*three.Face {
	return []*three.Face{
		three.NewFace(v1, v4, v3),
		three.NewFace(v1, v2, v4),
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
