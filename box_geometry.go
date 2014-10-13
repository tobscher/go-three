package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"github.com/go-gl/mathgl/mgl32"
)

type boxGeometry struct {
	bufferData    []float32
	buffer        gl.Buffer
	bufferLoaded  bool
	program       gl.Program
	programLoaded bool
	matrixID      gl.UniformLocation
}

func NewBoxGeometry(width, height, depth float32) *boxGeometry {
	pos := mgl32.Vec3{0, 0, 0}
	bufferData := make([]float32, 0)

	halfWidth := width / 2.0
	halfHeight := height / 2.0
	halfDepth := depth / 2.0

	// Bottom plane
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() - halfHeight, pos.Z() - halfDepth},
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() - halfHeight, pos.Z() - halfDepth},
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() - halfHeight, pos.Z() + halfDepth},
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() - halfHeight, pos.Z() + halfDepth},
	)...)

	// Side 1
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() - halfHeight, pos.Z() - halfDepth},
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() - halfHeight, pos.Z() - halfDepth},
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() + halfHeight, pos.Z() - halfDepth},
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() + halfHeight, pos.Z() - halfDepth},
	)...)

	// Side 2
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() - halfHeight, pos.Z() - halfDepth},
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() - halfHeight, pos.Z() + halfDepth},
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() + halfHeight, pos.Z() - halfDepth},
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() + halfHeight, pos.Z() + halfDepth},
	)...)

	// // Side 3
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() - halfHeight, pos.Z() - halfDepth},
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() - halfHeight, pos.Z() + halfDepth},
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() + halfHeight, pos.Z() - halfDepth},
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() + halfHeight, pos.Z() + halfDepth},
	)...)

	// // Side 4
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() - halfHeight, pos.Z() + halfDepth},
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() - halfHeight, pos.Z() + halfDepth},
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() + halfHeight, pos.Z() + halfDepth},
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() + halfHeight, pos.Z() + halfDepth},
	)...)

	// Top plane
	bufferData = append(bufferData, buildPlane(
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() + halfHeight, pos.Z() - halfDepth},
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() + halfHeight, pos.Z() - halfDepth},
		mgl32.Vec3{pos.X() - halfWidth, pos.Y() + halfHeight, pos.Z() + halfDepth},
		mgl32.Vec3{pos.X() + halfWidth, pos.Y() + halfHeight, pos.Z() + halfDepth},
	)...)
	return &boxGeometry{bufferData: bufferData, programLoaded: false, bufferLoaded: false}
}

func NewCubeGeometry(size float32) *boxGeometry {
	return NewBoxGeometry(size, size, size)
}

func buildPlane(v1, v2, v3, v4 mgl32.Vec3) []float32 {
	return []float32{
		v1.X(), v1.Y(), v1.Z(),
		v4.X(), v4.Y(), v4.Z(),
		v3.X(), v3.Y(), v3.Z(),
		v1.X(), v1.Y(), v1.Z(),
		v2.X(), v2.Y(), v2.Z(),
		v4.X(), v4.Y(), v4.Z(),
	}
}

func (bg *boxGeometry) Program() gl.Program {
	if !bg.programLoaded {
		vShader := glh.Shader{gl.VERTEX_SHADER, loadDataFile("triangle.v.glsl")}
		fShader := glh.Shader{gl.FRAGMENT_SHADER, loadDataFile("triangle.f.glsl")}
		bg.program = glh.NewProgram(vShader, fShader)
		bg.matrixID = bg.program.GetUniformLocation("MVP")

		bg.programLoaded = true
	}

	return bg.program
}

func (bg *boxGeometry) MatrixID() gl.UniformLocation {
	return bg.matrixID
}

func (bg *boxGeometry) Buffer() gl.Buffer {
	if !bg.bufferLoaded {
		bg.buffer = gl.GenBuffer()
		bg.buffer.Bind(gl.ARRAY_BUFFER)
		gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(bg.bufferData), bg.bufferData, gl.STATIC_DRAW)

		bg.bufferLoaded = true
	}

	return bg.buffer
}

func (bg *boxGeometry) VertexCount() int {
	// 6 sides * 2 (for each triangle)
	return 6 * 2 * 3
}
