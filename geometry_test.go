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
