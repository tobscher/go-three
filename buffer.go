package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"log"
)

type buffer struct {
	data     []float32
	glBuffer gl.Buffer
}

func NewBuffer(data []float32) buffer {
	log.Println("*** New Buffer generated ***")

	glBuffer := gl.GenBuffer()
	glBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(data), data, gl.STATIC_DRAW)

	buffer := buffer{glBuffer: glBuffer, data: data}

	return buffer
}

func (b *buffer) unload() {
	b.glBuffer.Delete()
}

func (b *buffer) update(data []float32) {
	log.Println("*** Buffer updated ***")

	b.data = data

	b.bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(data), data, gl.STATIC_DRAW)
}

func (b buffer) bind(enum gl.GLenum) {
	b.glBuffer.Bind(enum)
}

func (b buffer) vertexCount() int {
	return len(b.data) / 3
}
