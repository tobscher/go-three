package three

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

// Program is a GLSL shader program.
type Program struct {
	attributes map[string]Attribute
	glProgram  gl.Program
	Loaded     bool
	uniforms   map[string]*Uniform
}

// ProgramFeature type.
type ProgramFeature int

const (
	// COLOR feature
	COLOR ProgramFeature = 1 << iota
	// TEXTURE feature
	TEXTURE
)

var (
	featureDefinitions = map[ProgramFeature]string{
		COLOR:   "USE_COLOR",
		TEXTURE: "USE_TEXTURE",
	}
)

// NewProgram returns a new Program with attributes and uniforms collection initialised.
func NewProgram() *Program {
	return &Program{
		attributes: make(map[string]Attribute),
		uniforms:   make(map[string]*Uniform),
	}
}

// Load sets the given OpenGL program.
func (p *Program) Load(program gl.Program) {
	log.Println("*** Program loaded ***")

	p.glProgram = program
	p.Loaded = true
}

func (p *Program) unload() {
	for _, attribute := range p.attributes {
		attribute.buffer.Delete()
	}
	p.glProgram.Delete()
}

// Use makes this program current.
func (p Program) Use() {
	p.glProgram.Use()
}

// Unuse disables the current program.
func (p Program) Unuse() {
	p.glProgram.Unuse()
}

// MakeProgram loads a shader program for the given features.
// Features will be activated via pre-processor directives.
// e.g. #define USE_TEXTURE
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

#ifdef USE_TEXTURE
  layout(location = 1) in vec2 vertexUV;
  out vec2 UV;
#endif

layout(location = 2) in vec3 vertexNormal_modelspace;

uniform mat4 MVP;
uniform vec2 repeat;

void main() {
  gl_Position =  MVP * vec4(vertexPosition_modelspace,1);

#ifdef USE_TEXTURE
  UV = vertexUV * repeat;
#endif
}`, getShaderDefinitions(features))

	return formatted
}

func loadFragmentShader(features ProgramFeature) string {
	formatted := fmt.Sprintf(`
#version 330 core

%v

#ifdef USE_COLOR
  uniform vec3 diffuse;
#endif
out vec4 color;

#ifdef USE_TEXTURE
  in vec2 UV;
  uniform sampler2D textureSampler;
#endif

void main() {
#ifdef USE_COLOR
  color = vec4(diffuse, 1.0);
#endif

#ifdef USE_TEXTURE
  color = vec4(texture(textureSampler, UV).rgb, 1.0);
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
