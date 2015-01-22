package three

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl"
	glh "github.com/tobscher/glh"
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
	// SHADING_BASIC feature
	SHADING_BASIC
)

var (
	featureDefinitions = map[ProgramFeature]string{
		COLOR:         "USE_COLOR",
		TEXTURE:       "USE_TEXTURE",
		SHADING_BASIC: "USE_BASIC_SHADING",
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
	logger.Debug("Creating new shader program")

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

out vec3 Position_worldspace;
out vec3 Normal_cameraspace;
out vec3 EyeDirection_cameraspace;
out vec3 LightDirection_cameraspace;

uniform mat4 MVP;
uniform mat4 V;
uniform mat4 M;
uniform vec3 LightPosition_worldspace;
uniform vec2 repeat;

void main() {

  Position_worldspace = (M * vec4(vertexPosition_modelspace,1)).xyz;

  vec3 vertexPosition_cameraspace = ( V * M * vec4(vertexPosition_modelspace,1)).xyz;
  EyeDirection_cameraspace = vec3(0,0,0) - vertexPosition_cameraspace;

  vec3 LightPosition_cameraspace = ( V * vec4(LightPosition_worldspace,1)).xyz;
  LightDirection_cameraspace = LightPosition_cameraspace + EyeDirection_cameraspace;

  Normal_cameraspace = ( V * M * vec4(vertexNormal_modelspace,0)).xyz;

#ifdef USE_TEXTURE
  UV = vertexUV * repeat;
#endif

  gl_Position =  MVP * vec4(vertexPosition_modelspace,1);

}`, getShaderDefinitions(features))

	return formatted
}

func loadFragmentShader(features ProgramFeature) string {
	formatted := fmt.Sprintf(`
#version 330 core

%v

in vec3 Position_worldspace;
in vec3 Normal_cameraspace;
in vec3 EyeDirection_cameraspace;
in vec3 LightDirection_cameraspace;

#ifdef USE_COLOR
  uniform vec3 diffuse;
#endif
out vec3 color;

#ifdef USE_TEXTURE
  in vec2 UV;
  uniform sampler2D textureSampler;
#endif

uniform mat4 MV;
uniform vec3 LightPosition_worldspace;

void main() {
#ifdef USE_COLOR
  vec3 MaterialDiffuseColor = vec3(diffuse);
#endif

#ifdef USE_TEXTURE
  vec3 MaterialDiffuseColor = vec3(texture(textureSampler, UV).rgb);
#endif

vec3 LightColor = vec3(1,1,1);
float LightPower = 50.0f;

vec3 MaterialAmbientColor = vec3(0.1,0.1,0.1) * MaterialDiffuseColor;
vec3 MaterialSpecularColor = vec3(0.3,0.3,0.3);

float distance = length( LightPosition_worldspace - Position_worldspace );

vec3 n = normalize( Normal_cameraspace );
vec3 l = normalize( LightDirection_cameraspace );
float cosTheta = clamp( dot( n,l ), 0,1 );

vec3 E = normalize(EyeDirection_cameraspace);
vec3 R = reflect(-l,n);

float cosAlpha = clamp( dot( E,R ), 0,1 );

#ifdef USE_BASIC_SHADING
  color = MaterialAmbientColor +
          MaterialDiffuseColor * LightColor * LightPower * cosTheta / (distance*distance) +
          MaterialSpecularColor * LightColor * LightPower * pow(cosAlpha,5) / (distance*distance);
#else
  color = MaterialDiffuseColor;
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

	if hasFeature(features, SHADING_BASIC) {
		defines = append(defines, fmt.Sprintf("#define %v", featureDefinitions[SHADING_BASIC]))
	}

	return strings.Join(defines, "\n")
}

func hasFeature(features ProgramFeature, feature ProgramFeature) bool {
	return features&feature == feature
}
