package three

import "github.com/go-gl/mathgl/mgl32"

type MockGeometry struct {
}

func NewMockGeometry() *MockGeometry {
	return nil
}

func (m *MockGeometry) Vertices() []mgl32.Vec3 {
	return []mgl32.Vec3{
		mgl32.Vec3{0, 0, 0},
		mgl32.Vec3{-1, 1, 0},
		mgl32.Vec3{1, 1, 0},
	}
}

func (m *MockGeometry) Faces() []*Face {
	return []*Face{
		NewFace(1, 2, 3),
	}
}

func (m *MockGeometry) UVs() []mgl32.Vec2 {
	return []mgl32.Vec2{
		mgl32.Vec2{0, 0},
		mgl32.Vec2{1, 0},
		mgl32.Vec2{0, 1},
	}
}
