package three

type Mesh struct {
	Rotation Euler
}

func NewMesh(geometry BoxGeometry, material MeshBasicMaterial) Mesh {
	return Mesh{}
}
