package three

type Mesh struct {
	geometry triangleGeometry
	material meshBasicMaterial
}

func NewMesh(geometry triangleGeometry, material meshBasicMaterial) Mesh {
	return Mesh{geometry: geometry, material: material}
}
