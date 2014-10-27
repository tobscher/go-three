package three

import (
	gl "github.com/go-gl/gl"
)

// Attribute describes an attribute which is passed to
// a shader program.
type Attribute struct {
	index  int
	size   uint
	buffer *Buffer
}

// NewAttribute creates a new attribute for a shader program.
func NewAttribute(index int, size uint, buffer *Buffer) Attribute {
	return Attribute{index: index, size: size, buffer: buffer}
}

func (a *Attribute) enableFor() gl.AttribLocation {
	location := gl.AttribLocation(a.index)
	location.EnableArray()
	a.buffer.bind(gl.ARRAY_BUFFER)
	location.AttribPointer(a.size, gl.FLOAT, false, 0, nil)

	return location
}
