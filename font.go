package three

import (
	"os"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/tobscher/gltext"
)

type Font struct {
	texture *Texture
	font    *gltext.Font
}

func NewFont(path string, scale int32) (*Font, error) {
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
