package materials

import (
	three "github.com/tobscher/go-three"
)

// Basic describes a material with basic shading. No lights or shadows
// are considered.
type Basic struct {
	program *three.Program

	color     *three.Color
	texture   *three.Texture
	wireframe bool
}

// NewBasic creates a new Basic material.
func NewBasic() *Basic {
	material := Basic{}
	return &material
}

// SetColor sets a solid color for this material.
func (b *Basic) SetColor(color *three.Color) *Basic {
	b.color = color
	return b
}

// Color returns the set color for this material.
func (b Basic) Color() *three.Color {
	return b.color
}

// SetTexture sets the texture for this material.
func (b *Basic) SetTexture(texture *three.Texture) *Basic {
	b.texture = texture
	return b
}

// Texture returns the set texture for this material.
func (b Basic) Texture() *three.Texture {
	return b.texture
}

// SetWireframe enables or disables wireframes for the material.
func (b *Basic) SetWireframe(wireframe bool) *Basic {
	b.wireframe = wireframe
	return b
}

// Wireframe returns the current value for wireframe rendering.
func (b Basic) Wireframe() bool {
	return b.wireframe
}

// SetProgram stores the given program for further use (cache).
func (b *Basic) SetProgram(p *three.Program) {
	b.program = p
}

// Program returns the current program which is used for this material.
func (b *Basic) Program() *three.Program {
	return b.program
}
