package three

// Face stores indices to vertices and normals.
type Face struct {
	vertexIndices [3]uint16
	normalIndices [3]uint16
}

// NewFace returns a new triangle face.
func NewFace(a, b, c uint16) *Face {
	return &Face{
		vertexIndices: [3]uint16{a, b, c},
	}
}

// At returns the vertex index stored at the given position.
func (f *Face) At(i int) uint16 {
	return f.vertexIndices[i]
}

// NormalAt returns the normal index stored at the given position
func (f *Face) NormalAt(i int) uint16 {
	return f.normalIndices[i]
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

// AddNormal adds normal indices for this face.
func (f *Face) AddNormal(x, y, z uint16) {
	f.normalIndices[0] = x
	f.normalIndices[1] = y
	f.normalIndices[2] = z
}
