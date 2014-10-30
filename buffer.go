package three

import (
	"github.com/go-gl/gl"
	"log"
)

// Buffer consists of a buffer and it's underlying data.
type Buffer struct {
	data     []interface{}
	enum     gl.GLenum
	size     int
	glBuffer gl.Buffer
}

// NewBuffer creates a new OpenGL buffer and buffers the given data.
func NewBuffer(data []interface{}, enum gl.GLenum, size int) Buffer {
	log.Println("*** New Buffer generated ***", data)

	glBuffer := gl.GenBuffer()
	glBuffer.Bind(enum)
	gl.BufferData(enum, len(data)*size, data, gl.STATIC_DRAW)

	buffer := Buffer{glBuffer: glBuffer, data: data, enum: enum, size: size}

	return buffer
}

func (b *Buffer) unload() {
	b.glBuffer.Delete()
}

func (b *Buffer) update(data []interface{}) {
	log.Println("*** Buffer updated ***")

	b.data = data

	b.bind(b.enum)
	gl.BufferData(b.enum, len(data)*b.size, data, gl.STATIC_DRAW)
}

func (b *Buffer) bind(enum gl.GLenum) {
	b.glBuffer.Bind(enum)
}

func (b *Buffer) unbind(enum gl.GLenum) {
	b.glBuffer.Unbind(enum)
}
