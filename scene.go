package three

type scene struct {
	objects []*Mesh
}

func NewScene() scene {
	return scene{}
}

func (s *scene) Add(mesh *Mesh) {
	s.objects = append(s.objects, mesh)
}
