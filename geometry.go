package three

import "github.com/go-gl/mathgl/mgl32"

// Shape is the interface which defines the shape of a
// 3D object.
type Shape interface {
	Vertices() []mgl32.Vec3
	UVs() []mgl32.Vec2
	Normals() []mgl32.Vec3
	Faces() []*Face
	ArrayCount() int
}

// Geometry is a base struct with fields for Vertices and UVs.
type Geometry struct {
	vertices []mgl32.Vec3
	uvs      []mgl32.Vec2
	normals  []mgl32.Vec3
	faces    []*Face
}

// SetVertices stores the given vertices in an internal field.
func (g *Geometry) SetVertices(vertices []mgl32.Vec3) {
	g.vertices = vertices
}

// Vertices returns the vertices for the geometry.
func (g *Geometry) Vertices() []mgl32.Vec3 {
	return g.vertices
}

// SetNormals stores the given normals in an internal field.
func (g *Geometry) SetNormals(normals []mgl32.Vec3) {
	g.normals = normals
}

// Normals returns the normals for the geometry.
func (g *Geometry) Normals() []mgl32.Vec3 {
	return g.normals
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

// ArrayCount returns the number of elements to draw in the draw call.
// Faces * 3 if faces are used or number of vertices.
func (g *Geometry) ArrayCount() int {
	if len(g.faces) > 0 {
		return len(g.faces) * 3
	}

	return len(g.vertices)
}
