package three

type Mesh struct {
	buffers  map[string]buffer
	geometry Shape
	material Appearance

	Transform Transform
}

func NewMesh(geometry Shape, material Appearance) Mesh {
	m := Mesh{
		geometry:  geometry,
		material:  material,
		Transform: NewTransform(1),
	}

	return m
}
