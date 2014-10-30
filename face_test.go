package three

import (
	"testing"
)

func TestNewFace(t *testing.T) {
	face := NewFace(1, 2, 3)

	if face.A() != 1 {
		t.Errorf("Index A invalid. Expected 1 got %v", face.A())
	}

	if face.B() != 2 {
		t.Errorf("Index B invalid. Expected 2 got %v", face.B())
	}

	if face.C() != 3 {
		t.Errorf("Index C invalid. Expected 3 got %v", face.C())
	}
}
