package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

func GenerateBuffer(bufferData []float32) gl.Buffer {
	buffer := gl.GenBuffer()
	buffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(bufferData), bufferData, gl.STATIC_DRAW)

	return buffer
}
