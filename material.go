package three

type Material interface {
	generateColorBuffer(int) []float32
	generateUvBuffer(int) []float32

	Color() *Color
	ColorsDirty() bool
	SetColorsDirty(bool)

	Texture() *Texture
	TextureDirty() bool
	SetTextureDirty(bool)

	Program(*Mesh) Program
	Wireframe() bool
}
