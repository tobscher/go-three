package three

import ()

type meshBasicMaterial struct {
	attributes []Attribute
	program    Program

	color     Color
	texture   Texture
	wireframe bool
}

func NewMeshBasicMaterial() *meshBasicMaterial {
	material := meshBasicMaterial{}
	return &material
}

func (m meshBasicMaterial) Program() Program {
	if !m.program.loaded {
		m.program.load(MakeProgram(COLOR))
	}

	return m.program
}

func (m meshBasicMaterial) Attributes() []Attribute {
	return m.attributes
}

func (m meshBasicMaterial) SetColor(color Color) meshBasicMaterial {
	m.color = color
	m.attributes = append(m.attributes, NewAttribute(1, getColorBuffer(36, color)))
	return m
}

func (m meshBasicMaterial) Color() Color {
	return m.color
}

func (m meshBasicMaterial) SetTexture(texture Texture) meshBasicMaterial {
	m.texture = texture
	return m
}

func (m meshBasicMaterial) Texture() Texture {
	return m.texture
}

func (m meshBasicMaterial) SetWireframe(wireframe bool) meshBasicMaterial {
	m.wireframe = wireframe
	return m
}

func (m meshBasicMaterial) Wireframe() bool {
	return m.wireframe
}

func getColorBuffer(verticesCount int, color Color) buffer {
	bufferData := make([]float32, 0, verticesCount*3)
	for i := 0; i < verticesCount; i++ {
		bufferData = append(bufferData, color.R(), color.G(), color.B())
	}

	buffer := NewBuffer(bufferData)

	return buffer
}
