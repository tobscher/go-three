package three

import (
	gl "github.com/go-gl/gl"
)

type Index struct {
	buffer *Buffer
}

func NewIndex(buffer *Buffer) *Index {
	return &Index{buffer: buffer}
}

func (i *Index) enable() {
	i.buffer.bind(gl.ELEMENT_ARRAY_BUFFER)
}

func (i *Index) disable() {
	i.buffer.unbind(gl.ELEMENT_ARRAY_BUFFER)
}
