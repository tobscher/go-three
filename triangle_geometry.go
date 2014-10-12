package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

type TriangleGeometry struct {
	bufferData    [9]float32
	buffer        gl.Buffer
	bufferLoaded  bool
	program       gl.Program
	programLoaded bool
	matrixID      gl.UniformLocation
}

func NewTriangleGeometry(bufferData [9]float32) TriangleGeometry {
	return TriangleGeometry{bufferData: bufferData, programLoaded: false, bufferLoaded: false}
}

func (tg *TriangleGeometry) Program() gl.Program {
	if !tg.programLoaded {
		vShader := glh.Shader{gl.VERTEX_SHADER, loadDataFile("triangle.v.glsl")}
		fShader := glh.Shader{gl.FRAGMENT_SHADER, loadDataFile("triangle.f.glsl")}
		tg.program = glh.NewProgram(vShader, fShader)
		tg.matrixID = tg.program.GetUniformLocation("MVP")

		tg.programLoaded = true
	}

	return tg.program
}

func (tg *TriangleGeometry) MatrixID() gl.UniformLocation {
	return tg.matrixID
}

func (tg *TriangleGeometry) Buffer() gl.Buffer {
	if !tg.bufferLoaded {
		tg.buffer = gl.GenBuffer()
		tg.buffer.Bind(gl.ARRAY_BUFFER)
		gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(tg.bufferData), &tg.bufferData, gl.STATIC_DRAW)

		tg.bufferLoaded = true
	}

	return tg.buffer
}

func (tg *TriangleGeometry) vertexCount() int {
	return 3
}
