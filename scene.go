package three

// Scene represents a tree-like structure of 3D objects.
type Scene struct {
	objects []*Mesh
}

// NewScene returns a new Scene.
func NewScene() *Scene {
	return &Scene{}
}

// Add adds the given mesh to the scene tree.
func (s *Scene) Add(mesh *Mesh) {
	s.objects = append(s.objects, mesh)
}
