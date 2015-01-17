package three

import "github.com/go-gl/gl"

// Object3D described an object in 3D space.
type Object3D struct {
	vertexBuffer gl.Buffer
	uvBuffer     gl.Buffer
	index        *Index

	transform *Transform
}

// NewObject3D returns a new Object3D. It initializes the transform
// of the object.
func NewObject3D() Object3D {
	o := Object3D{
		transform: NewTransform(),
	}

	return o
}

// Index returns the FBO index.
func (o *Object3D) Index() *Index {
	return o.index
}

// Transform returns the transform object (translation, rotation, scale) for this object.
func (o *Object3D) Transform() *Transform {
	return o.transform
}

// UVBuffer returns the buffer object for this objects uvs.
func (o *Object3D) UVBuffer() gl.Buffer {
	return o.uvBuffer
}

// VertexBuffer returns the buffer object for this objects vertices.
func (o *Object3D) VertexBuffer() gl.Buffer {
	return o.vertexBuffer
}

func newUvBuffer(geometry Shape) gl.Buffer {
	result := []float32{}

	for _, uv := range geometry.UVs() {
		result = append(result, uv.X(), uv.Y())
	}

	// Invert V because we're using a compressed texture
	for i := 1; i < len(result); i += 2 {
		result[i] = 1.0 - result[i]
	}

	glBuffer := gl.GenBuffer()
	glBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(result)*2*4, result, gl.STATIC_DRAW)

	return glBuffer
}

func newVertexBuffer(geometry Shape) gl.Buffer {
	result := []float32{}

	for _, vertex := range geometry.Vertices() {
		result = append(result, vertex.X(), vertex.Y(), vertex.Z())
	}

	glBuffer := gl.GenBuffer()
	glBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(result)*3*4, result, gl.STATIC_DRAW)

	return glBuffer
}

func generateIndex(geometry Shape) *Index {
	data := []uint16{}

	for _, f := range geometry.Faces() {
		data = append(data, f.A(), f.B(), f.C())
	}

	index := NewIndex(data)
	return index
}
