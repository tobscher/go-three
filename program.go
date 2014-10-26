package three

import (
	"fmt"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"log"
	"strings"
)

// Knows about attributes and uniforms
type Program struct {
	attributes map[string]Attribute
	glProgram  gl.Program
	Loaded     bool
	matrixID   gl.UniformLocation
}

type ProgramFeature int

const (
	COLOR ProgramFeature = 1 << iota
	TEXTURE
)

var (
	featureDefinitions = map[ProgramFeature]string{
		COLOR:   "USE_COLOR",
		TEXTURE: "USE_TEXTURE",
	}
)

func NewProgram() *Program {
	return &Program{
		attributes: make(map[string]Attribute),
	}
}

func (p *Program) Load(program gl.Program) {
	log.Println("*** Program loaded ***")

	p.glProgram = program
	p.matrixID = p.glProgram.GetUniformLocation("MVP")
	p.Loaded = true
}

func (p *Program) unload() {
	p.glProgram.Delete()
}

func (p Program) use() {
	p.glProgram.Use()
}

func (p Program) MatrixID() gl.UniformLocation {
	return p.matrixID
}

func MakeProgram(features ProgramFeature) gl.Program {
	vertSource := loadVertexShader(features)
	fragSource := loadFragmentShader(features)

	return glh.NewProgram(
		glh.Shader{Type: gl.VERTEX_SHADER, Program: string(vertSource)},
		glh.Shader{Type: gl.FRAGMENT_SHADER, Program: string(fragSource)},
	)
}

func loadVertexShader(features ProgramFeature) string {
	formatted := fmt.Sprintf(`
#version 330 core

%v

layout(location = 0) in vec3 vertexPosition_modelspace;

#ifdef USE_COLOR
  layout(location = 1) in vec3 vertexColor;
  out vec3 fragmentColor;
#endif

#ifdef USE_TEXTURE
  layout(location = 1) in vec2 vertexUV;
  out vec2 UV;
#endif

uniform mat4 MVP;

void main() {
  gl_Position =  MVP * vec4(vertexPosition_modelspace,1);

#ifdef USE_COLOR
  fragmentColor = vertexColor;
#endif

#ifdef USE_TEXTURE
  UV = vertexUV;
#endif
}`, getShaderDefinitions(features))

	return formatted
}

func loadFragmentShader(features ProgramFeature) string {
	formatted := fmt.Sprintf(`
#version 330 core

%v

#ifdef USE_COLOR
  in vec3 fragmentColor;
#endif
out vec3 color;

#ifdef USE_TEXTURE
  in vec2 UV;
  uniform sampler2D textureSampler;
#endif

void main() {
#ifdef USE_COLOR
  color = fragmentColor;
#endif

#ifdef USE_TEXTURE
  color = texture(textureSampler, UV).rgb;
#endif
}`, getShaderDefinitions(features))

	return formatted
}

func getShaderDefinitions(features ProgramFeature) string {
	defines := []string{}

	if hasFeature(features, COLOR) {
		defines = append(defines, fmt.Sprintf("#define %v", featureDefinitions[COLOR]))
	}

	if hasFeature(features, TEXTURE) {
		defines = append(defines, fmt.Sprintf("#define %v", featureDefinitions[TEXTURE]))
	}

	return strings.Join(defines, "\n")
}

func hasFeature(features ProgramFeature, feature ProgramFeature) bool {
	return features&feature == feature
}
