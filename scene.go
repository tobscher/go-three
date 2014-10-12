package three

type Scene struct {
	objects []*Mesh
}

func NewScene() Scene {
	return Scene{}
}

func (s *Scene) Add(mesh *Mesh) {
	s.objects = append(s.objects, mesh)
}
