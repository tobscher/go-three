package materials

import (
	three "github.com/tobscher/go-three"
)

type Basic struct {
	program *three.Program

	color     *three.Color
	texture   *three.Texture
	wireframe bool
}

func NewBasic() *Basic {
	material := Basic{}
	return &material
}

func (b *Basic) SetColor(color *three.Color) *Basic {
	b.color = color
	return b
}

func (b Basic) Color() *three.Color {
	return b.color
}

func (b *Basic) SetTexture(texture *three.Texture) *Basic {
	b.texture = texture
	return b
}

func (b Basic) Texture() *three.Texture {
	return b.texture
}

func (b *Basic) SetWireframe(wireframe bool) *Basic {
	b.wireframe = wireframe
	return b
}

func (b Basic) Wireframe() bool {
	return b.wireframe
}

func (b *Basic) SetProgram(p *three.Program) {
	b.program = p
}

func (b *Basic) Program() *three.Program {
	return b.program
}
