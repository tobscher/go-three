package three

import "github.com/go-gl/gl"

// Text defines a 2D text geometry and its appearance.
type Text struct {
	vertexBuffer gl.Buffer
	uvBuffer     gl.Buffer

	material Texter
	geometry *TextGeometry
}

// NewText creates a new 2D text object that can be added to
// the scene.
//
// NOTE: This method overrides the texture for the given texture
// to the texture from the font.
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

// Geometry returns the geometry object for this
// text object.
func (t *Text) Geometry() *TextGeometry {
	return t.geometry
}

// Material returns the appearance defintion for this
// text object.
func (t *Text) Material() Texter {
	return t.material
}

// SetText updates the text that is drawn.
// This will update the underlying vertex and uv buffer.
// BUG(tobscher) Vertex and UV buffers are deleted and re-created; not updated.
func (t *Text) SetText(text string) {
	t.geometry.Text = text
	t.geometry.UpdateVertices(text)

	t.vertexBuffer.Delete()
	t.uvBuffer.Delete()

	t.vertexBuffer = newTextVertexBuffer(t.geometry)
	t.uvBuffer = newUvBuffer(t.geometry.UVs, true)
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
