package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Uniform struct {
	location gl.UniformLocation
}

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
	}
}
