package three

type meshBasicMaterial struct {
	program Program

	colorsDirty  bool
	color        *Color
	textureDirty bool
	texture      *texture
	wireframe    bool
}

func NewMeshBasicMaterial() *meshBasicMaterial {
	material := meshBasicMaterial{}
	return &material
}

func (m *meshBasicMaterial) Program(mesh *Mesh) Program {
	if !m.program.loaded {
		if m.color != nil {
			m.program.load(MakeProgram(COLOR))
			m.program.attributes["color"] = NewAttribute(1, 3, &mesh.colorBuffer)
		}

		if m.texture != nil {
			m.program.load(MakeProgram(TEXTURE))
			m.program.attributes["texture"] = NewAttribute(1, 2, &mesh.uvBuffer)
		}
	}

	return m.program
}

func (m *meshBasicMaterial) SetColor(color *Color) *meshBasicMaterial {
	m.color = color
	m.colorsDirty = true
	return m
}

func (m meshBasicMaterial) Color() *Color {
	return m.color
}

func (m meshBasicMaterial) ColorsDirty() bool {
	return m.colorsDirty
}

func (m *meshBasicMaterial) SetColorsDirty(dirty bool) {
	m.colorsDirty = dirty
}

func (m *meshBasicMaterial) SetTexture(texture *texture) *meshBasicMaterial {
	m.textureDirty = true
	m.texture = texture
	return m
}

func (m meshBasicMaterial) Texture() *texture {
	return m.texture
}

func (m meshBasicMaterial) TextureDirty() bool {
	return m.textureDirty
}

func (m *meshBasicMaterial) SetTextureDirty(dirty bool) {
	m.textureDirty = dirty
}

func (m *meshBasicMaterial) SetWireframe(wireframe bool) *meshBasicMaterial {
	m.wireframe = wireframe
	return m
}

func (m meshBasicMaterial) Wireframe() bool {
	return m.wireframe
}

func (m *meshBasicMaterial) generateColorBuffer(vertexCount int) []float32 {
	bufferData := make([]float32, 0, vertexCount*3)
	for i := 0; i < vertexCount; i++ {
		bufferData = append(bufferData, m.color.R(), m.color.G(), m.color.B())
	}

	return bufferData
}

func (m *meshBasicMaterial) generateUvBuffer(vertexCount int) []float32 {
	bufferData := make([]float32, 0, vertexCount*2)
	for i := 0; i < 6; i++ {
		bufferData = append(bufferData,
			1, 0,
			0, 1,
			1, 1,

			1, 0,
			0, 0,
			0, 1,
		)
	}

	// bufferData := []float32{
	// 	// Visible
	// 	1, 0, // bottom right
	// 	0, 1, // top left
	// 	1, 1, // top right

	// 	1, 0, // bottom right
	// 	0, 0, // bottom left
	// 	0, 1, // top left
	// }

	// Invert V because we're using a compressed texture
	for i := 1; i < len(bufferData); i += 2 {
		bufferData[i] = 1.0 - bufferData[i]
	}

	return bufferData
}
