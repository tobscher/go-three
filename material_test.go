package three

type MockMaterial struct {
	program *Program
}

func NewMockMaterial() *MockMaterial {
	return &MockMaterial{}
}

func (m *MockMaterial) Program() *Program {
	return m.program
}

func (m *MockMaterial) SetProgram(program *Program) {
	m.program = program
}
