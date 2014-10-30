package three

import (
	gl "github.com/go-gl/gl"
	glh "github.com/tobscher/glh"
)

type Index struct {
	glBuffer gl.Buffer
	count    int
}

func NewIndex(data []uint16) *Index {
	glBuffer := gl.GenBuffer()
	glBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(data)*int(glh.Sizeof(gl.UNSIGNED_SHORT)), data, gl.STATIC_DRAW)

	return &Index{glBuffer: glBuffer, count: len(data)}
}

func (i *Index) enable() {
	i.glBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
}

func (i *Index) disable() {
	i.glBuffer.Unbind(gl.ELEMENT_ARRAY_BUFFER)
}
