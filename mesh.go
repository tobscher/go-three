package three

type Mesh struct {
	geometry Geometry
	material meshBasicMaterial
}

func NewMesh(geometry Geometry, material meshBasicMaterial) Mesh {
	return Mesh{geometry: geometry, material: material}
}
