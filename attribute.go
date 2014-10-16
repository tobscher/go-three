package three

import (
	gl "github.com/go-gl/gl"
)

type Attribute struct {
	index  int
	buffer *buffer
}

func NewAttribute(index int, buffer *buffer) Attribute {
	return Attribute{index: index, buffer: buffer}
}

// Bug(tobscher) Buffer data per feature should probably be cached
// and not computed every frame as it doesn't change unless
// the material changes
func (a *Attribute) enableFor(m *Mesh) gl.AttribLocation {
	location := gl.AttribLocation(a.index)
	location.EnableArray()
	a.buffer.bind(gl.ARRAY_BUFFER)
	location.AttribPointer(3, gl.FLOAT, false, 0, nil)

	return location
}
