package three

import (
	"github.com/go-gl/gl"
)

// Texture represents a graphic that can be rendered on any geometry.
type Texture struct {
	glTexture gl.Texture
}

// NewTexture returns a new DDS texture loaded from the given path.
func NewTexture(path string) (*Texture, error) {
	t, err := TextureFromDDS(path)
	if err != nil {
		return nil, err
	}

	return &Texture{glTexture: t}, nil
}

// Unload deallocates the opengl texture object.
func (t *Texture) Unload() {
	t.glTexture.Delete()
}
