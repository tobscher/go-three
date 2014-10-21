package three

type MockMaterial struct {
}

func NewMockMaterial() *MockMaterial {
	return nil
}

func (m *MockMaterial) Color() *Color {
	return nil
}

func (m *MockMaterial) ColorsDirty() bool {
	return false
}

func (m *MockMaterial) SetColorsDirty(dirty bool) {
}

func (m *MockMaterial) Texture() *texture {
	return nil
}

func (m *MockMaterial) TextureDirty() bool {
	return false
}

func (m *MockMaterial) SetTextureDirty(dirty bool) {
}

func (m *MockMaterial) Program(mesh *Mesh) Program {
	return Program{}
}

func (m *MockMaterial) Wireframe() bool {
	return false
}

func (m *MockMaterial) generateColorBuffer(n int) []float32 {
	return []float32{
		1, 0, 0,
		1, 0, 0,
		1, 0, 0,
	}
}

func (m *MockMaterial) generateUvBuffer(n int) []float32 {
	return []float32{
		1, 0, 0,
		1, 0, 0,
		1, 0, 0,
	}
}
