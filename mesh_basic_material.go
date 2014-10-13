package three

import (
	"github.com/go-gl/gl"
)

type meshBasicMaterial struct {
	buffer       gl.Buffer
	bufferLoaded bool

	color     Color
	wireframe bool
}

func NewMeshBasicMaterial() *meshBasicMaterial {
	material := meshBasicMaterial{bufferLoaded: false}
	return &material
}

func (m meshBasicMaterial) SetColor(color Color) meshBasicMaterial {
	m.color = color
	return m
}

func (m meshBasicMaterial) Color() Color {
	return m.color
}

func (m meshBasicMaterial) SetWireframe(wireframe bool) meshBasicMaterial {
	m.wireframe = wireframe
	return m
}

func (m meshBasicMaterial) Wireframe() bool {
	return m.wireframe
}

func (m *meshBasicMaterial) Buffer(verticesCount int) gl.Buffer {
	if !m.bufferLoaded {
		bufferData := make([]float32, 0, verticesCount*3)
		for i := 0; i < verticesCount; i++ {
			bufferData = append(bufferData, m.Color().R(), m.Color().G(), m.Color().B())
		}

		m.buffer = gl.GenBuffer()
		m.buffer.Bind(gl.ARRAY_BUFFER)
		gl.BufferData(gl.ARRAY_BUFFER, len(bufferData)*4, bufferData, gl.STATIC_DRAW)
		m.bufferLoaded = true
	}

	return m.buffer
}
