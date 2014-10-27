package three

// BasicMaterial describes a material with basic shading. No lights or shadows
// are considered.
type BasicMaterial struct {
	program *Program

	color     *Color
	texture   *Texture
	wireframe bool
}

// NewBasic creates a new Basic material.
func NewBasicMaterial() *BasicMaterial {
	material := BasicMaterial{}
	return &material
}

// SetColor sets a solid color for this material.
func (b *BasicMaterial) SetColor(color *Color) *BasicMaterial {
	b.color = color
	return b
}

// Color returns the set color for this material.
func (b BasicMaterial) Color() *Color {
	return b.color
}

// SetTexture sets the texture for this material.
func (b *BasicMaterial) SetTexture(texture *Texture) *BasicMaterial {
	b.texture = texture
	return b
}

// Texture returns the set texture for this material.
func (b BasicMaterial) Texture() *Texture {
	return b.texture
}

// SetWireframe enables or disables wireframes for the material.
func (b *BasicMaterial) SetWireframe(wireframe bool) *BasicMaterial {
	b.wireframe = wireframe
	return b
}

// Wireframe returns the current value for wireframe rendering.
func (b BasicMaterial) Wireframe() bool {
	return b.wireframe
}

// SetProgram stores the given program for further use (cache).
func (b *BasicMaterial) SetProgram(p *Program) {
	b.program = p
}

// Program returns the current program which is used for this material.
func (b *BasicMaterial) Program() *Program {
	return b.program
}
