package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"log"
)

// Buffer consists of a buffer and it's underlying data.
type Buffer struct {
	data     []float32
	glBuffer gl.Buffer
}

// NewBuffer creates a new OpenGL buffer and buffers the given data.
func NewBuffer(data []float32) Buffer {
	log.Println("*** New Buffer generated ***")

	glBuffer := gl.GenBuffer()
	glBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(data), data, gl.STATIC_DRAW)

	buffer := Buffer{glBuffer: glBuffer, data: data}

	return buffer
}

func (b *Buffer) unload() {
	b.glBuffer.Delete()
}

func (b *Buffer) update(data []float32) {
	log.Println("*** Buffer updated ***")

	b.data = data

	b.bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(data), data, gl.STATIC_DRAW)
}

func (b Buffer) bind(enum gl.GLenum) {
	b.glBuffer.Bind(enum)
}
