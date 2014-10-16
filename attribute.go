package three

import (
	gl "github.com/go-gl/gl"
)

type Attribute struct {
	index   int
	feature ProgramFeature
	buffer  buffer
}

func NewAttribute(index int, feature ProgramFeature) Attribute {
	return Attribute{index: index, feature: feature}
}

// Bug(tobscher) Buffer data per feature should probably be cached
// and not computed every frame as it doesn't change unless
// the material changes
func (a Attribute) enableFor(m *Mesh) gl.AttribLocation {
	bufferData := m.material.BufferDataFor(a.feature, m.geometry)
	a.buffer = NewBuffer(bufferData)
	a.buffer.load()

	location := gl.AttribLocation(a.index)
	location.EnableArray()
	a.buffer.bind(gl.ARRAY_BUFFER)
	location.AttribPointer(3, gl.FLOAT, false, 0, nil)

	return location
}
