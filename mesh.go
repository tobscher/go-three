package three

// Mesh is a representation of a 3D object. It consists
// of a geometry and a material. Meshes can be transformed
// in 3D space.
type Mesh struct {
	index    *Index
	geometry Shape
	material Appearance

	Transform Transform
}

// NewMesh creates a new mesh for the given geometry (shape) and material (appearance).
func NewMesh(geometry Shape, material Appearance) Mesh {
	m := Mesh{
		geometry:  geometry,
		material:  material,
		Transform: NewTransform(1),
	}

	return m
}
