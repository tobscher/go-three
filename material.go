package three

import (
	gl "github.com/go-gl/gl"
)

type Material interface {
	Color() Color
	Buffer(verticesCount int) gl.Buffer
	Wireframe() bool
}
