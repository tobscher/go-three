package three

func NewMockMesh() Mesh {
	geometry := NewMockGeometry()
	material := NewMockMaterial()

	return Mesh{geometry: geometry, material: material}
}
