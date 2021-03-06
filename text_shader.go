package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

// MakeTextShader creates a new shader program
// for text rendering purposes.
func MakeTextShader() gl.Program {
	vertSource := loadTextVertexShader()
	fragSource := loadTextFragmentShader()

	return glh.NewProgram(
		glh.Shader{Type: gl.VERTEX_SHADER, Program: string(vertSource)},
		glh.Shader{Type: gl.FRAGMENT_SHADER, Program: string(fragSource)},
	)
}

func loadTextVertexShader() string {
	vertex := `
#version 330

// Input vertex data, different for all executions of this shader.
layout(location = 0) in vec2 vertexPosition_screenspace;
layout(location = 1) in vec2 vertexUV;

// Output data ; will be interpolated for each fragment.
out vec2 UV;

void main(){

  // Output position of the vertex, in clip space
  // map [0..800][0..600] to [-1..1][-1..1]
  vec2 vertexPosition_homoneneousspace = vertexPosition_screenspace - vec2(400,300); // [0..800][0..600] -> [-400..400][-300..300]
  vertexPosition_homoneneousspace /= vec2(400,300);
  gl_Position =  vec4(vertexPosition_homoneneousspace,0,1);

  // UV of the vertex. No special space for this one.
  UV = vertexUV;
}`

	return vertex
}

func loadTextFragmentShader() string {
	fragment := `
#version 330

// Interpolated values from the vertex shaders
in vec2 UV;

// Ouput data
out vec4 color;

// Values that stay constant for the whole mesh.
uniform vec3 diffuse;
uniform sampler2D textureSampler;

void main(){
  color = vec4(diffuse, 1.0) * texture(textureSampler, UV);
}
`

	return fragment
}
