package three

type Material interface {
	generateColorBuffer(int) []float32

	ColorsDirty() bool
	SetColorsDirty(bool)

	Program(*Mesh) Program
	Wireframe() bool
}
