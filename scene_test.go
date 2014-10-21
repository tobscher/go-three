package three

import (
	"testing"
)

func TestAddToScene(t *testing.T) {
	scene := NewScene()
	mesh := NewMockMesh()
	scene.Add(&mesh)

	expected := 1
	actual := len(scene.objects)

	if actual != expected {
		t.Errorf("Object was not added to the scene: Scene elements expected %v got %v", expected, actual)
	}

	if scene.objects[0] != &mesh {
		t.Errorf("Object is not the element that was added.")
	}
}
