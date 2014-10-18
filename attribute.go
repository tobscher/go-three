package three

import (
	gl "github.com/go-gl/gl"
)

type Attribute struct {
	index  int
	size   uint
	buffer *buffer
}

func NewAttribute(index int, size uint, buffer *buffer) Attribute {
	return Attribute{index: index, size: size, buffer: buffer}
}

// Bug(tobscher) Buffer data per feature should probably be cached
// and not computed every frame as it doesn't change unless
// the material changes
func (a *Attribute) enableFor(m *Mesh) gl.AttribLocation {
	location := gl.AttribLocation(a.index)
	location.EnableArray()
	a.buffer.bind(gl.ARRAY_BUFFER)
	location.AttribPointer(a.size, gl.FLOAT, false, 0, nil)

	return location
}
