package three

type Mesh struct {
	vertexBuffer buffer
	colorBuffer  buffer
	uvBuffer     buffer

	geometry Geometry
	material Material

	Transform Transform
}

func NewMesh(geometry Geometry, material Material) Mesh {
	vertexBuffer := NewBuffer(geometry.generateVertexBuffer())

	m := Mesh{
		vertexBuffer: vertexBuffer,
		geometry:     geometry,
		material:     material,
		Transform:    NewTransform(1),
	}

	if material.Color() != nil {
		colorBuffer := NewBuffer(material.generateColorBuffer(vertexBuffer.vertexCount()))
		m.colorBuffer = colorBuffer
	}

	if material.Texture() != nil {
		uvBuffer := NewBuffer(material.generateUvBuffer(vertexBuffer.vertexCount()))
		m.uvBuffer = uvBuffer
	}

	return m
}

func (m *Mesh) ColorBuffer() buffer {
	if m.material.ColorsDirty() {
		m.colorBuffer.update(m.material.generateColorBuffer(m.vertexBuffer.vertexCount()))
		m.material.SetColorsDirty(false)
	}

	return m.colorBuffer
}

func (m *Mesh) UvBuffer() buffer {
	if m.material.TextureDirty() {
		m.uvBuffer.update(m.material.generateUvBuffer(m.vertexBuffer.vertexCount()))
		m.material.SetTextureDirty(false)
	}

	return m.uvBuffer
}
