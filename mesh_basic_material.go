package three

import (
	"github.com/go-gl/gl"
)

type MeshBasicMaterial struct {
	buffer       gl.Buffer
	bufferLoaded bool

	Color Color
}

func NewMeshBasicMaterial(color Color) MeshBasicMaterial {
	material := MeshBasicMaterial{bufferLoaded: false, Color: color}
	return material
}

func (m *MeshBasicMaterial) Buffer(verticesCount int) gl.Buffer {
	if !m.bufferLoaded {
		bufferData := make([]float32, 0, verticesCount*3)
		for i := 0; i < verticesCount; i++ {
			bufferData = append(bufferData, m.Color.R(), m.Color.G(), m.Color.B())
		}

		m.buffer = gl.GenBuffer()
		m.buffer.Bind(gl.ARRAY_BUFFER)
		gl.BufferData(gl.ARRAY_BUFFER, len(bufferData)*4, bufferData, gl.STATIC_DRAW)
		m.bufferLoaded = true
	}

	return m.buffer
}
