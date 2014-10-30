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

	halfWidth := width / 2.0
	halfHeight := height / 2.0
	halfDepth := depth / 2.0

	vertices := []mgl32.Vec3{
		// front
		mgl32.Vec3{-halfWidth, -halfHeight, halfDepth},
		mgl32.Vec3{halfWidth, -halfHeight, halfDepth},
		mgl32.Vec3{halfWidth, halfHeight, halfDepth},
		mgl32.Vec3{-halfWidth, halfHeight, halfDepth},
		// top
		mgl32.Vec3{-halfWidth, halfHeight, halfDepth},
		mgl32.Vec3{halfWidth, halfHeight, halfDepth},
		mgl32.Vec3{halfWidth, halfHeight, -halfDepth},
		mgl32.Vec3{-halfWidth, halfHeight, -halfDepth},
		// back
		mgl32.Vec3{halfWidth, -halfHeight, -halfDepth},
		mgl32.Vec3{-halfWidth, -halfHeight, -halfDepth},
		mgl32.Vec3{-halfWidth, halfHeight, -halfDepth},
		mgl32.Vec3{halfWidth, halfHeight, -halfDepth},
		//bottom
		mgl32.Vec3{-halfWidth, -halfHeight, -halfDepth},
		mgl32.Vec3{halfWidth, -halfHeight, -halfDepth},
		mgl32.Vec3{halfWidth, -halfHeight, halfDepth},
		mgl32.Vec3{-halfWidth, -halfHeight, halfDepth},
		// left
		mgl32.Vec3{-halfWidth, -halfHeight, -halfDepth},
		mgl32.Vec3{-halfWidth, -halfHeight, halfDepth},
		mgl32.Vec3{-halfWidth, halfHeight, halfDepth},
		mgl32.Vec3{-halfWidth, halfHeight, -halfDepth},
		// right
		mgl32.Vec3{halfWidth, -halfHeight, halfDepth},
		mgl32.Vec3{halfWidth, -halfHeight, -halfDepth},
		mgl32.Vec3{halfWidth, halfHeight, -halfDepth},
		mgl32.Vec3{halfWidth, halfHeight, halfDepth},
	}

	var uvs []mgl32.Vec2

	for i := 0; i < 6; i++ {
		uvs = append(uvs,
			mgl32.Vec2{0.0, 0.0},
			mgl32.Vec2{1.0, 0.0},
			mgl32.Vec2{1.0, 1.0},
			mgl32.Vec2{0.0, 1.0},
		)
	}

	faces := []*three.Face{
		// front
		three.NewFace(0, 1, 2),
		three.NewFace(2, 3, 0),
		// top
		three.NewFace(4, 5, 6),
		three.NewFace(6, 7, 4),
		// back
		three.NewFace(8, 9, 10),
		three.NewFace(10, 11, 8),
		// bottom
		three.NewFace(12, 13, 14),
		three.NewFace(14, 15, 12),
		// left
		three.NewFace(16, 17, 18),
		three.NewFace(18, 19, 16),
		// right
		three.NewFace(20, 21, 22),
		three.NewFace(22, 23, 20),
	}

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
