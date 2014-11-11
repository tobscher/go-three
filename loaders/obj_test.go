package loaders

import "testing"

func TestLoadFromObjVertices(t *testing.T) {
	geometry, err := LoadFromObj("../test/fixtures/box.obj")
	if err != nil {
		t.Fatal(err)
	}

	expected := 8
	actual := len(geometry.Vertices())

	if actual != expected {
		t.Errorf("Expected %v vertices got %v.", expected, actual)
	}
}

func TestLoadFromObjFaces(t *testing.T) {
	geometry, err := LoadFromObj("../test/fixtures/box.obj")
	if err != nil {
		t.Fatal(err)
	}

	expected := 12
	actual := len(geometry.Faces())

	if actual != expected {
		t.Errorf("Expected %v faces got %v.", expected, actual)
	}
}
