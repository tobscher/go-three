package three

import "github.com/go-gl/gl"

type Text struct {
	vertexBuffer gl.Buffer
	uvBuffer     gl.Buffer

	material Texter
	geometry *TextGeometry
}

func NewText(geometry *TextGeometry, material Texter) *Text {
	material.SetTexture(geometry.Font.texture)

	t := Text{
		material: material,
		geometry: geometry,
	}

	t.vertexBuffer = newTextVertexBuffer(t.geometry)
	t.uvBuffer = newUvBuffer(t.geometry.UVs, true)

	return &t
}

func newTextVertexBuffer(geometry *TextGeometry) gl.Buffer {
	result := []float32{}

	vertices := geometry.Vertices

	for _, vertex := range vertices {
		result = append(result, vertex.X(), vertex.Y())
	}

	glBuffer := gl.GenBuffer()
	glBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(result)*2*4, result, gl.STATIC_DRAW)

	return glBuffer
}

func (t *Text) Geometry() *TextGeometry {
	return t.geometry
}

func (t *Text) Material() Texter {
	return t.material
}
