package three

import (
	gl "github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Geometry interface {
	updateBuffer(mgl32.Vec3)
	Program() gl.Program
	MatrixID() gl.UniformLocation
	Buffer() gl.Buffer
	VertexCount() int
}
