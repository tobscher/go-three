package three

import (
	"fmt"
	"os"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/tobscher/gltext"
)

// Font defines a font with its texture
type Font struct {
	texture *Texture
	font    *gltext.Font
}

// NewFont loads a font file from path at a given scale.
// The font will be rendered to a texture which is later used
// for drawing.
func NewFont(path string, scale int32) (*Font, error) {
	logger.Info(fmt.Sprintf("Loading font from: %v", path))

	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer fd.Close()

	f, err := gltext.LoadTruetype(fd, scale, 32, 127, gltext.LeftToRight)
	if err != nil {
		return nil, err
	}

	texture := Texture{
		glTexture: f.Texture,
		WrapS:     ClampToEdge,
		WrapT:     ClampToEdge,
		Repeat:    mgl32.Vec2{1, 1},
	}

	return &Font{
		font:    f,
		texture: &texture,
	}, nil
}
