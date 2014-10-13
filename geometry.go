package three

import (
	gl "github.com/go-gl/gl"
)

type Geometry interface {
	Program() gl.Program
	MatrixID() gl.UniformLocation
	Buffer() gl.Buffer
	VertexCount() int
}
