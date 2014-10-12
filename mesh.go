package three

type Mesh struct {
	geometry TriangleGeometry
	material MeshBasicMaterial
}

func NewMesh(geometry TriangleGeometry, material MeshBasicMaterial) Mesh {
	return Mesh{geometry: geometry, material: material}
}
