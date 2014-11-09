package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Wrapping describes an enum for texture wrapping modes.
// Possible modes are: ClampToEdge, Repeat or MirroredRepeat
type Wrapping int

const (
	// ClampToEdge extends texture to the nearest edge, texture is not repeated
	ClampToEdge Wrapping = iota
	// Repeat repeats the texture
	Repeat
	// MirroredRepeat repeats texture mirrored
	MirroredRepeat
)

// Texture represents a graphic that can be rendered on any geometry.
type Texture struct {
	glTexture gl.Texture
	WrapS     Wrapping
	WrapT     Wrapping
	Repeat    mgl32.Vec2
}

// NewTexture returns a new DDS texture loaded from the given path.
// WrapS and WrapT are defaultd to ClampToEdge.
// Repeat is set to 1,1 (no repeat)
func NewTexture(path string) (*Texture, error) {
	t, err := TextureFromDDS(path)
	if err != nil {
		return nil, err
	}

	texture := Texture{
		glTexture: t,
		WrapS:     ClampToEdge,
		WrapT:     ClampToEdge,
		Repeat:    mgl32.Vec2{1, 1},
	}

	return &texture, nil
}

// Bind binds this texture.
func (t *Texture) Bind() {
	t.glTexture.Bind(gl.TEXTURE_2D)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, t.getWrapping(t.WrapS))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, t.getWrapping(t.WrapT))
}

// Unbind unbinds this texture.
func (t *Texture) Unbind() {
	t.glTexture.Unbind(gl.TEXTURE_2D)
}

// Unload deallocates the opengl texture object.
func (t *Texture) Unload() {
	t.glTexture.Delete()
}

func (t *Texture) getWrapping(wr Wrapping) int {
	var result gl.GLenum
	switch wr {
	case Repeat:
		result = gl.REPEAT
	case ClampToEdge:
		result = gl.CLAMP_TO_EDGE
	case MirroredRepeat:
		result = gl.MIRRORED_REPEAT
	}

	return int(result)
}
