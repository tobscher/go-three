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

const (
	SHADER_VERSION = "#version 330 core"
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
	formatted := fmt.Sprintf(`%v

%v

%v

%v

void main() {
  %v
}`, SHADER_VERSION, loadVertexIns(features), loadVertexOuts(features), loadVertexUniforms(features), loadVertexFeatures(features))
	return formatted
}

func loadFragmentShader(features ProgramFeature) string {
	formatted := fmt.Sprintf(`%v

%v
%v

%v

void main() {
  %v
}`, SHADER_VERSION, loadFragmentIns(features), loadFragmentOuts(features), loadFragmentUniforms(features), loadFragmentFeatures(features))
	return formatted
}

// Fragment
func loadFragmentIns(features ProgramFeature) string {
	ins := []string{}

	if features&COLOR == COLOR {
		ins = append(ins, "in vec3 fragmentColor;")
	}

	if features&TEXTURE == TEXTURE {
		ins = append(ins, "in vec2 UV;")
	}
	return strings.Join(ins, "\n")
}

func loadFragmentOuts(features ProgramFeature) string {
	return "out vec3 color;"
}

func loadFragmentUniforms(features ProgramFeature) string {
	if features&TEXTURE == TEXTURE {
		return "uniform sampler2D textureSampler;"
	}

	return "// No uniforms"
}

func loadFragmentFeatures(features ProgramFeature) string {
	if features&COLOR == COLOR {
		return "color = fragmentColor;"
	}

	if features&TEXTURE == TEXTURE {
		return "color = texture(textureSampler, UV).rgb;"
	}

	return ""
}

// Vertex
func loadVertexIns(features ProgramFeature) string {
	ins := []string{
		"layout(location = 0) in vec3 vertexPosition_modelspace;",
	}

	if features&COLOR == COLOR {
		ins = append(ins, "layout(location = 1) in vec3 vertexColor;")
	}

	if features&TEXTURE == TEXTURE {
		ins = append(ins, "layout(location = 1) in vec2 vertexUV;")
	}

	return strings.Join(ins, "\n")
}

func loadVertexOuts(features ProgramFeature) string {
	outs := []string{}

	if features&COLOR == COLOR {
		outs = append(outs, "out vec3 fragmentColor;")
	}

	if features&TEXTURE == TEXTURE {
		outs = append(outs, "out vec2 UV;")
	}

	return strings.Join(outs, "\n")
}

func loadVertexUniforms(features ProgramFeature) string {
	uniforms := []string{
		"uniform mat4 MVP;",
	}

	return strings.Join(uniforms, "\n")
}

func loadVertexFeatures(features ProgramFeature) string {
	result := []string{
		"gl_Position =  MVP * vec4(vertexPosition_modelspace,1);",
	}

	if features&COLOR == COLOR {
		result = append(result, "  fragmentColor = vertexColor;")
	}

	if features&TEXTURE == TEXTURE {
		result = append(result, "  UV = vertexUV;")
	}

	return strings.Join(result, "\n")
}
