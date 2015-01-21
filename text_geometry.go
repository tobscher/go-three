package three

import "github.com/go-gl/mathgl/mgl32"

// TextGeometry defines the geometry of 2D text.
type TextGeometry struct {
	Vertices []mgl32.Vec2
	UVs      []mgl32.Vec2

	Text     string
	Position mgl32.Vec2
	Size     float32
	Font     *Font
}

// NewTextGeometry creates a new 2D text geometry for the given text.
//
// NOTE: Y-Axis for position is inverted!
func NewTextGeometry(text string, position mgl32.Vec2, size float32, font *Font) *TextGeometry {
	vertices, uvs := createTextVertices(text, position, size, font)

	geometry := TextGeometry{
		Vertices: vertices,
		UVs:      uvs,

		Text:     text,
		Size:     size,
		Position: position,
		Font:     font,
	}

	return &geometry
}

func (t *TextGeometry) updateVertices(text string) {
	vertices, uvs := createTextVertices(text, t.Position, t.Size, t.Font)

	t.Vertices = vertices
	t.UVs = uvs
}

func createTextVertices(text string, position mgl32.Vec2, size float32, font *Font) (vertices []mgl32.Vec2, uvs []mgl32.Vec2) {
	x := position.X()
	y := position.Y()

	for c, char := range text {
		i := float32(c)

		upLeft := mgl32.Vec2{x + i*size, y + size}
		upRight := mgl32.Vec2{x + i*size + size, y + size}
		downRight := mgl32.Vec2{x + i*size + size, y}
		downLeft := mgl32.Vec2{x + i*size, y}

		vertices = append(vertices, upLeft, downLeft, upRight)
		vertices = append(vertices, downRight, upRight, downLeft)

		glyph := font.font.Glyphs().Find(string(char))
		fullWidth := float32(font.font.Width)
		fullHeight := float32(font.font.Height)
		width := float32(glyph.Width)
		height := float32(glyph.Height)
		x := float32(glyph.X)
		y := float32(glyph.Y)

		uvX := x / fullWidth
		uvY := (fullHeight - y) / fullHeight
		uvWidth := width / fullWidth
		uvHeight := height / fullHeight

		uvUpLeft := mgl32.Vec2{uvX, uvY}
		uvUpRight := mgl32.Vec2{uvX + uvWidth, uvY}
		uvDownRight := mgl32.Vec2{uvX + uvWidth, uvY - uvHeight}
		uvDownLeft := mgl32.Vec2{uvX, uvY - uvHeight}

		uvs = append(uvs, uvUpLeft, uvDownLeft, uvUpRight)
		uvs = append(uvs, uvDownRight, uvUpRight, uvDownLeft)
	}

	return
}
