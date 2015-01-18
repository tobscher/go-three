package three

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"
)

type TextGeometry struct {
	Vertices []mgl32.Vec2
	UVs      []mgl32.Vec2

	Text     string
	Position mgl32.Vec2
	Font     *Font
}

func NewTextGeometry(text string, position mgl32.Vec2, size float32, font *Font) *TextGeometry {
	vertices := []mgl32.Vec2{}
	uvs := []mgl32.Vec2{}

	x := position.X()
	y := position.Y()

	log.Println(text)

	for c, char := range text {
		i := float32(c)

		up_left := mgl32.Vec2{x + i*size, y + size}
		up_right := mgl32.Vec2{x + i*size + size, y + size}
		down_right := mgl32.Vec2{x + i*size + size, y}
		down_left := mgl32.Vec2{x + i*size, y}

		vertices = append(vertices, up_left, down_left, up_right)
		vertices = append(vertices, down_right, up_right, down_left)

		glyph := font.font.Glyphs().Find(string(char))
		fullWidth := float32(font.font.Width)
		fullHeight := float32(font.font.Height)
		width := float32(glyph.Width)
		height := float32(glyph.Height)
		x := float32(glyph.X)
		y := float32(glyph.Y)

		// log.Printf("%+v", glyph)
		// log.Printf("Full width: %v\n", fullWidth)
		// log.Printf("Full height: %v\n", fullHeight)
		// log.Printf("width: %v\n", width)
		// log.Printf("height: %v\n", height)
		// log.Printf("x: %v\n", x)
		// log.Printf("y: %v\n", y)

		var uv_x float32 = x / fullWidth
		var uv_y float32 = (fullHeight - y) / fullHeight
		var uv_width float32 = width / fullWidth
		var uv_height float32 = height / fullHeight

		uv_up_left := mgl32.Vec2{uv_x, uv_y}
		uv_up_right := mgl32.Vec2{uv_x + uv_width, uv_y}
		uv_down_right := mgl32.Vec2{uv_x + uv_width, uv_y - uv_height}
		uv_down_left := mgl32.Vec2{uv_x, uv_y - uv_height}

		uvs = append(uvs, uv_up_left, uv_down_left, uv_up_right)
		uvs = append(uvs, uv_down_right, uv_up_right, uv_down_left)
	}

	// log.Printf("Vertices: %+v", vertices)
	// log.Printf("UVs: %+v", uvs)

	geometry := TextGeometry{
		Vertices: vertices,
		UVs:      uvs,

		Text:     text,
		Position: position,
		Font:     font,
	}

	return &geometry
}
