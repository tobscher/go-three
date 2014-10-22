package three

// Material is the interface which defines the appearance of 3D objects.
// For example: solid color, textured, etc.
//
// Bug(tobscher) This interface is to big and needs breaking down into smaller pieces.
// For example: Colored, Textured, etc.
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
