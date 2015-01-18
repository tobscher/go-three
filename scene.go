package three

import "github.com/go-gl/gl"

// SceneObject is an interface that describes an object
// that can be added to the scene graph.
type SceneObject interface {
	Geometry() Shape
	Material() Appearance
	Transform() *Transform
	Mode() gl.GLenum

	Index() *Index
	VertexBuffer() gl.Buffer
	UVBuffer() gl.Buffer
	NormalBuffer() gl.Buffer
}

// Scene represents a tree-like structure (graph) of 3D objects.
type Scene struct {
	objects []SceneObject
	texts   []*Text
}

// NewScene returns a new Scene.
func NewScene() *Scene {
	return &Scene{}
}

// Add adds the given scene object to the scene tree.
func (s *Scene) Add(object SceneObject) {
	s.objects = append(s.objects, object)
}

func (s *Scene) AddText(text *Text) {
	s.texts = append(s.texts, text)
}
