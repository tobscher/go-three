package three

import (
	"github.com/go-gl/gl"
)

type texture struct {
	glTexture gl.Texture
}

func NewTexture(path string) (*texture, error) {
	t, err := TextureFromDDS(path)
	if err != nil {
		return nil, err
	}

	return &texture{glTexture: t}, nil
}

func (t *texture) unload() {
	t.glTexture.Delete()
}
