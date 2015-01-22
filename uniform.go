package three

import (
	"fmt"

	"github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tobscher/gltext"
)

// Uniform describes a uniform value which is passed
// to the shader program.
type Uniform struct {
	identifier string
	location   gl.UniformLocation
}

// NewUniform creates a new uniform using the given identifier to lookup
// the uniform position in the shader program.
func NewUniform(program *Program, identifier string) *Uniform {
	logger.Debug(fmt.Sprintf("Uniform: %v", identifier))

	uniform := Uniform{
		identifier: identifier,
		location:   program.glProgram.GetUniformLocation(identifier),
	}
	return &uniform
}

func (u *Uniform) apply(value interface{}) {
	logger.Trace(fmt.Sprintf("Applying uniform: %v = %v", u.identifier, value))

	switch t := value.(type) {
	case mgl32.Mat4:
		u.location.UniformMatrix4fv(false, t)
	case *Color:
		u.location.Uniform3fv(1, t.Float())
	case *Texture:
		u.location.Uniform1i(0)
	case mgl32.Vec2:
		u.location.Uniform2f(t[0], t[1])
	case mgl32.Vec3:
		u.location.Uniform3f(t[0], t[1], t[2])
	case gltext.Glyph:
		u.location.Uniform4f(float32(t.X), float32(t.Y), float32(t.Width), float32(t.Height))
	}
}
