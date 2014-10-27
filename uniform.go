package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Uniform describes a uniform value which is passed
// to the shader program.
type Uniform struct {
	location gl.UniformLocation
}

// NewUniform creates a new uniform using the given identifier to lookup
// the uniform position in the shader program.
func NewUniform(program *Program, identifier string) *Uniform {
	uniform := Uniform{
		program.glProgram.GetUniformLocation(identifier),
	}
	return &uniform
}

func (u *Uniform) apply(value interface{}) {
	switch t := value.(type) {
	case mgl32.Mat4:
		u.location.UniformMatrix4fv(false, t)
	case *Color:
		u.location.Uniform3fv(1, t.Float())
	}
}
