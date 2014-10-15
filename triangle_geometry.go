package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"github.com/go-gl/mathgl/mgl32"
)

type triangleGeometry struct {
	bufferData    []float32
	buffer        gl.Buffer
	bufferLoaded  bool
	program       gl.Program
	programLoaded bool
	matrixID      gl.UniformLocation
}

func NewTriangleGeometry(p1, p2, p3 mgl32.Vec3) *triangleGeometry {
	bufferData := []float32{
		p1.X(), p1.Y(), p1.Z(),
		p2.X(), p2.Y(), p2.Z(),
		p3.X(), p3.Y(), p3.Z(),
	}
	return &triangleGeometry{bufferData: bufferData, programLoaded: false, bufferLoaded: false}
}

func (tg *triangleGeometry) Program() gl.Program {
	if !tg.programLoaded {
		tg.program = MakeProgram("triangle.v.glsl", "triangle.f.glsl")
		tg.matrixID = tg.program.GetUniformLocation("MVP")

		tg.programLoaded = true
	}

	return tg.program
}

func (tg *triangleGeometry) MatrixID() gl.UniformLocation {
	return tg.matrixID
}

func (tg *triangleGeometry) Buffer() gl.Buffer {
	if !tg.bufferLoaded {
		tg.buffer = gl.GenBuffer()
		tg.buffer.Bind(gl.ARRAY_BUFFER)
		gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(tg.bufferData), tg.bufferData, gl.STATIC_DRAW)

		tg.bufferLoaded = true
	}

	return tg.buffer
}

func (tg *triangleGeometry) VertexCount() int {
	return 3
}
