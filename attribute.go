package three

import (
	gl "github.com/go-gl/gl"
)

// Attribute describes an attribute which is passed to
// a shader program.
type Attribute struct {
	index    int
	size     uint
	buffer   *Buffer
	location gl.AttribLocation
}

// NewAttribute creates a new attribute for a shader program.
func NewAttribute(index int, size uint, buffer *Buffer) Attribute {
	return Attribute{index: index, size: size, buffer: buffer}
}

func (a *Attribute) enable() {
	location := gl.AttribLocation(a.index)
	location.EnableArray()

	a.location = location
}

func (a *Attribute) pointer() {
	a.location.AttribPointer(a.size, gl.FLOAT, false, 0, nil)
}

func (a *Attribute) bindBuffer() {
	a.buffer.bind(gl.ARRAY_BUFFER)
}

func (a *Attribute) unbindBuffer() {
	a.buffer.unbind(gl.ARRAY_BUFFER)
}

func (a *Attribute) disable() {
	a.location.DisableArray()
}
