package three

import (
	gl "github.com/go-gl/gl"
)

type Attribute struct {
	index  int
	buffer buffer
}

func NewAttribute(index int, buffer buffer) Attribute {
	return Attribute{index: index, buffer: buffer}
}

func (a Attribute) enableFor(geometry Geometry) gl.AttribLocation {
	if !a.buffer.loaded {
		a.buffer.load()
	}

	location := gl.AttribLocation(a.index)
	location.EnableArray()
	a.buffer.bind(gl.ARRAY_BUFFER)
	location.AttribPointer(3, gl.FLOAT, false, 0, nil)

	return location
}
