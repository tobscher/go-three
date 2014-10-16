package three

type meshBasicMaterial struct {
	program Program

	colorsDirty bool
	color       Color
	texture     Texture
	wireframe   bool
}

func NewMeshBasicMaterial() *meshBasicMaterial {
	material := meshBasicMaterial{}
	return &material
}

func (m *meshBasicMaterial) Program(mesh *Mesh) Program {
	if !m.program.loaded {
		m.program.load(MakeProgram(COLOR))
		m.program.attributes["color"] = NewAttribute(1, &mesh.colorBuffer)
	}

	return m.program
}

func (m *meshBasicMaterial) SetColor(color Color) *meshBasicMaterial {
	m.color = color
	m.colorsDirty = true
	return m
}

func (m meshBasicMaterial) Color() Color {
	return m.color
}

func (m meshBasicMaterial) ColorsDirty() bool {
	return m.colorsDirty
}

func (m *meshBasicMaterial) SetColorsDirty(dirty bool) {
	m.colorsDirty = dirty
}

func (m meshBasicMaterial) SetTexture(texture Texture) meshBasicMaterial {
	m.texture = texture
	return m
}

func (m meshBasicMaterial) Texture() Texture {
	return m.texture
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
