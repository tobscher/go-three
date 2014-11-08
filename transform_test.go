package three

import (
	"math"
	"testing"

	"github.com/go-gl/mathgl/mgl32"
)

func TestNewTransform(t *testing.T) {
	transform := NewTransform()
	defaultPosition := mgl32.Vec3{0, 0, 0}
	defaultRotation := mgl32.Vec3{0, 0, 0}
	defaultQuaternion := mgl32.QuatIdent()
	defaultUp := mgl32.Vec3{0, 1, 0}
	defaultRight := mgl32.Vec3{1, 0, 0}
	defaultForward := mgl32.Vec3{0, 0, -1}
	defaultMatrix := mgl32.Ident4()

	if transform.position != defaultPosition {
		t.Errorf("Default position was not %v got %v", defaultPosition, transform.position)
	}

	if transform.rotation != defaultRotation {
		t.Errorf("Default rotation was not %v got %v", defaultRotation, transform.rotation)
	}

	if transform.quaternion != defaultQuaternion {
		t.Errorf("Default quaternion was not identity, got: %v", transform.quaternion)
	}

	if transform.Up != defaultUp {
		t.Errorf("Expected up to be %v got %v", defaultUp, transform.Up)
	}

	if transform.Right != defaultRight {
		t.Errorf("Expected right to be %v got %v", defaultRight, transform.Right)
	}

	if transform.Forward != defaultForward {
		t.Errorf("Expected forward to be %v got %v", defaultForward, transform.Forward)
	}

	if transform.matrix != defaultMatrix {
		t.Errorf("Expected default matrix to be identity matrix.")
	}
}

func TestTransformSetPosition(t *testing.T) {
	transform := NewTransform()

	expectedPosition := mgl32.Vec3{5, 10, 3.5}
	transform.SetPosition(expectedPosition.X(), expectedPosition.Y(), expectedPosition.Z())

	if transform.position != expectedPosition {
		t.Errorf("Expected position to be %v got %v", expectedPosition, transform.position)
	}

	if transform.matrix[12] != expectedPosition.X() ||
		transform.matrix[13] != expectedPosition.Y() ||
		transform.matrix[14] != expectedPosition.Z() {
		t.Error("Matrix was not updated correctly after translation.")
	}
}

func TestTransformTranslateX(t *testing.T) {
	transform := NewTransform()

	var translation float32 = 2.5
	transform.TranslateX(translation)

	if transform.position.X() != translation {
		t.Errorf("Expected position:x to be %v got %v", translation, transform.position.X())
	}

	if transform.matrix[12] != translation {
		t.Error("Matrix was not updated correctly after translation.")
	}
}

func TestTransformTranslateY(t *testing.T) {
	transform := NewTransform()

	var translation float32 = 3.75
	transform.TranslateY(translation)

	if transform.position.Y() != translation {
		t.Errorf("Expected position:y to be %v got %v", translation, transform.position.Y())
	}

	if transform.matrix[13] != translation {
		t.Error("Matrix was not updated correctly after translation.")
	}
}

func TestTransformTranslateZ(t *testing.T) {
	transform := NewTransform()

	var translation float32 = -13.37
	transform.TranslateZ(translation)

	if transform.position.Z() != translation {
		t.Errorf("Expected position:z to be %v got %v", translation, transform.position.Z())
	}

	if transform.matrix[14] != translation {
		t.Error("Matrix was not updated correctly after translation.")
	}
}

func TestTransformTranslate(t *testing.T) {
	transform := NewTransform()
	transform.SetPosition(10, 10, 10)

	vector := mgl32.Vec3{-5, 6, 3}
	expected := mgl32.Vec3{5, 16, 13}
	transform.Translate(vector)

	if transform.position != expected {
		t.Errorf("Expected position to be %v got %v", expected, transform.position)
	}

	if transform.matrix[12] != expected.X() ||
		transform.matrix[13] != expected.Y() ||
		transform.matrix[14] != expected.Z() {
		t.Error("Matrix was not updated correctly after translation.")
	}
}

func TestTransformScale(t *testing.T) {
	transform := NewTransform()
	scale := mgl32.Vec3{2, 2, 2}

	transform.Scale(scale.X(), scale.Y(), scale.Z())

	if transform.scale != scale {
		t.Errorf("Expected scale to be %v got %v", scale, transform.scale)
	}

	if transform.matrix[0] != scale.X() ||
		transform.matrix[5] != scale.Y() ||
		transform.matrix[10] != scale.Z() {
		t.Error("Matrix was not updated correctly after translation.")
	}
}

func TestTransformRotateX(t *testing.T) {
	transform := NewTransform()
	var rotation float32 = math.Pi / 2

	transform.RotateX(rotation)

	if transform.rotation.X() != rotation {
		t.Errorf("Expected rotation:x to be %v got %v", rotation, transform.rotation.X())
	}
}
