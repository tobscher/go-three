package three

import gl "github.com/go-gl/gl"

// Mesh is a representation of a 3D object. It consists
// of a geometry and a material. Meshes can be transformed
// in 3D space.
type Mesh struct {
	vertexBuffer gl.Buffer
	normalBuffer gl.Buffer
	uvBuffer     gl.Buffer
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
		Transform: NewTransform(),
	}

	m.vertexBuffer = newVertexBuffer(geometry)

	if len(geometry.UVs()) > 0 {
		m.uvBuffer = newUvBuffer(geometry)
	}

	if len(geometry.Normals()) > 0 {
		m.normalBuffer = newNormalBuffer(geometry)
	}
	m.index = generateIndex(geometry)

	return m
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

func newNormalBuffer(geometry Shape) gl.Buffer {
	result := []float32{}

	for _, normal := range geometry.Normals() {
		result = append(result, normal.X(), normal.Y(), normal.Z())
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
