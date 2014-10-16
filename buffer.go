package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

type buffer struct {
	data     []float32
	glBuffer gl.Buffer
	loaded   bool
}

func NewBuffer(bufferData []float32) buffer {
	return buffer{data: bufferData, loaded: false}
}

func (b *buffer) load() {
	b.glBuffer = gl.GenBuffer()
	b.glBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(b.data), b.data, gl.STATIC_DRAW)
	b.loaded = true
}

func (b buffer) bind(enum gl.GLenum) {
	b.glBuffer.Bind(enum)
}

func (b buffer) vertexCount() int {
	return len(b.data) / 3
}
