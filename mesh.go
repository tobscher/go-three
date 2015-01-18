package three

import "github.com/go-gl/gl"

// Mesh is a representation of a 3D object. It consists
// of a geometry and a material. Meshes can be transformed
// in 3D space.
type Mesh struct {
	Object3D

	geometry Shape
	material Appearance
}

// NewMesh creates a new mesh for the given geometry (shape) and material (appearance).
func NewMesh(geometry Shape, material Appearance) *Mesh {
	m := Mesh{
		Object3D: NewObject3D(),
		geometry: geometry,
		material: material,
	}

	m.vertexBuffer = newVertexBuffer(geometry)

	if len(geometry.UVs()) > 0 {
		m.uvBuffer = newUvBuffer(geometry.UVs(), true)
	}

	if len(geometry.Normals()) > 0 {
		m.normalBuffer = newNormalBuffer(geometry)
	}
	m.index = generateIndex(geometry)

	return &m
}

// Geometry returns an object which describes the shape of the mesh.
func (m *Mesh) Geometry() Shape {
	return m.geometry
}

// Material returns an objects which described the appaerance of the mesh.
func (m *Mesh) Material() Appearance {
	return m.material
}

// Mode returns an enum which is used to define the type of primitive to render.
func (m *Mesh) Mode() gl.GLenum {
	return gl.TRIANGLES
}
