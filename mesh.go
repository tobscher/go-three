package three

type Mesh struct {
	Rotation Euler
}

func NewMesh(geometry TriangleGeometry, material MeshBasicMaterial) Mesh {
	return Mesh{}
}
