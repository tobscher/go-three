package three

// Face stores indices to vertices.
type Face struct {
	vertexIndices [3]uint16
}

// NewFace returns a new triangle face.
func NewFace(a, b, c uint16) *Face {
	return &Face{
		vertexIndices: [3]uint16{a, b, c},
	}
}

// A returns the index of the first vertex in the triangle.
func (f *Face) A() uint16 {
	return f.vertexIndices[0]
}

// B returns the index of the second vertex in the triangle.
func (f *Face) B() uint16 {
	return f.vertexIndices[1]
}

// C returns the index of the third vertex in the triangle.
func (f *Face) C() uint16 {
	return f.vertexIndices[2]
}
