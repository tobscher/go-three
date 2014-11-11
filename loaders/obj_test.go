package loaders

import "testing"

func TestLoadFromObj(t *testing.T) {
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
