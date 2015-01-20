package three

// TextMaterial describes a material for 2D text.
type TextMaterial struct {
	program *Program

	color     *Color
	texture   *Texture
	wireframe bool
}

// NewTextMaterial creates a new Basic material.
func NewTextMaterial() *TextMaterial {
	material := TextMaterial{}
	return &material
}

// SetColor sets a solid color for this material.
func (b *TextMaterial) SetColor(color *Color) *TextMaterial {
	b.color = color
	return b
}

// Color returns the set color for this material.
func (b TextMaterial) Color() *Color {
	return b.color
}

// SetTexture sets the texture for this material.
func (b *TextMaterial) SetTexture(texture *Texture) *TextMaterial {
	b.texture = texture
	return b
}

// Texture returns the set texture for this material.
func (b TextMaterial) Texture() *Texture {
	return b.texture
}

// SetWireframe enables or disables wireframes for the material.
func (b *TextMaterial) SetWireframe(wireframe bool) *TextMaterial {
	b.wireframe = wireframe
	return b
}

// Wireframe returns the current value for wireframe rendering.
func (b TextMaterial) Wireframe() bool {
	return b.wireframe
}

// SetProgram stores the given program for further use (cache).
func (b *TextMaterial) SetProgram(p *Program) {
	b.program = p
}

// Program returns the current program which is used for this material.
func (b *TextMaterial) Program() *Program {
	return b.program
}
