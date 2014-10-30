package three

import (
	gl "github.com/go-gl/gl"
	"log"
)

// Mesh is a representation of a 3D object. It consists
// of a geometry and a material. Meshes can be transformed
// in 3D space.
type Mesh struct {
	vertexBuffer gl.Buffer
	index        *Index
	geometry     Shape
	material     Appearance

	Transform Transform
}

// NewMesh creates a new mesh for the given geometry (shape) and material (appearance).
func NewMesh(geometry Shape, material Appearance) Mesh {
	m := Mesh{
		geometry:  geometry,
		material:  material,
		Transform: NewTransform(1),
	}

	m.vertexBuffer = newVertexBuffer(geometry)
	m.index = generateIndex(geometry)

	return m
}

func newVertexBuffer(geometry Shape) gl.Buffer {
	result := []float32{}

	for _, vertex := range geometry.Vertices() {
		result = append(result, vertex.X(), vertex.Y(), vertex.Z())
	}

	glBuffer := gl.GenBuffer()
	glBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(result)*3*4, result, gl.STATIC_DRAW)

	log.Println("Vertices: ", result)

	return glBuffer
}

func generateIndex(geometry Shape) *Index {
	data := []uint16{}

	for _, f := range geometry.Faces() {
		data = append(data, f.A(), f.B(), f.C())
	}

	log.Println("Elements: ", data)

	index := NewIndex(data)
	return index
}
