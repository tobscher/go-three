package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"io/ioutil"
)

func MakeProgram(vertexShaderName, fragmentShaderName string) gl.Program {
	vertSource, err := ioutil.ReadFile("shaders/" + vertexShaderName)
	if err != nil {
		panic(err)
	}

	fragSource, err := ioutil.ReadFile("shaders/" + fragmentShaderName)
	if err != nil {
		panic(err)
	}

	return glh.NewProgram(glh.Shader{gl.VERTEX_SHADER, string(vertSource)},
		glh.Shader{gl.FRAGMENT_SHADER, string(fragSource)})
}
