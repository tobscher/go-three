package three

import "github.com/go-gl/gl"

type SceneObject interface {
	Program() gl.Program
	MatrixID() gl.UniformLocation
	Buffer() gl.Buffer
	vertexCount() int
}

type Scene struct {
	objects []*TriangleGeometry
}

func NewScene() Scene {
	return Scene{}
}

func (s *Scene) Add(sceneObject *TriangleGeometry) {
	s.objects = append(s.objects, sceneObject)
}
