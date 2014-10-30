package three

import "github.com/go-gl/mathgl/mgl32"

// Shape is the interface which defines the shape of a
// 3D object.
type Shape interface {
	Vertices() []mgl32.Vec3
	Faces() []*Face
	UVs() []mgl32.Vec2
}

// Geometry is a base struct with fields for Vertices and UVs.
type Geometry struct {
	vertices []mgl32.Vec3
	faces    []*Face
	uvs      []mgl32.Vec2
}

// SetVertices stores the given vertices in an internal field.
func (g *Geometry) SetVertices(vertices []mgl32.Vec3) {
	g.vertices = vertices
}

// Vertices returns the vertices for the geometry.
func (g *Geometry) Vertices() []mgl32.Vec3 {
	return g.vertices
}

// SetFaces stores the given faces in an internal field.
func (g *Geometry) SetFaces(faces []*Face) {
	g.faces = faces
}

// Faces returns the triangle faces for the geometry.
func (g *Geometry) Faces() []*Face {
	return g.faces
}

// SetUVs stores the given uv mappings in an internal field.
func (g *Geometry) SetUVs(uvs []mgl32.Vec2) {
	g.uvs = uvs
}

// UVs returns the uv mappings for the geometry.
func (g *Geometry) UVs() []mgl32.Vec2 {
	return g.uvs
}
