package three

type MockGeometry struct {
}

func NewMockGeometry() *MockGeometry {
	return nil
}

func (m *MockGeometry) generateVertexBuffer() []float32 {
	return []float32{
		0, 0, 0,
		-1, 1, 0,
		1, 1, 0,
	}
}
