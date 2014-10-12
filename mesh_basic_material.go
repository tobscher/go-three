package three

import "github.com/go-gl/gl"

type MeshBasicMaterial struct {
	r            float32
	g            float32
	b            float32
	buffer       gl.Buffer
	bufferData   [9]float32
	bufferLoaded bool
}

func NewMeshBasicMaterial(r, g, b float32) MeshBasicMaterial {
	return MeshBasicMaterial{bufferLoaded: false, r: r, g: g, b: b}
}

func (m *MeshBasicMaterial) Buffer() gl.Buffer {
	if !m.bufferLoaded {
		m.bufferData = [...]float32{
			m.r, m.g, m.b,
			m.r, m.g, m.b,
			m.r, m.g, m.b,
		}
		m.buffer = gl.GenBuffer()
		m.buffer.Bind(gl.ARRAY_BUFFER)
		gl.BufferData(gl.ARRAY_BUFFER, len(m.bufferData)*4, &m.bufferData, gl.STATIC_DRAW)
		m.bufferLoaded = true
	}

	return m.buffer
}
