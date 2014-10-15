package three

import (
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"testing"
)

func TestFragmentShaderWithSolidColor(t *testing.T) {
	expected := `#version 330 core

in vec3 fragmentColor;
out vec3 color;

// No uniforms

void main() {
  color = fragmentColor;
}`
	result := loadFragmentShader(COLOR)

	if result != expected {
		t.Errorf("Fragment shader invalid.\n\n***Expected***:\n%v\n\n***got***:\n%v", expected, result)
	}
}

func TestVertexShaderWithSolidColor(t *testing.T) {
	expected := `#version 330 core

layout(location = 0) in vec3 vertexPosition_modelspace;
layout(location = 1) in vec3 vertexColor;

out vec3 fragmentColor;

uniform mat4 MVP;

void main() {
  gl_Position =  MVP * vec4(vertexPosition_modelspace,1);
  fragmentColor = vertexColor;
}`
	result := loadVertexShader(COLOR)

	if result != expected {
		t.Errorf("Vertex shader invalid.\n\n***Expected***:\n%v\n\n***got***:\n%v", expected, result)
	}
}

func TestFragmentShaderWithTexture(t *testing.T) {
	expected := `#version 330 core

in vec2 UV;
out vec3 color;

uniform sampler2D textureSampler;

void main() {
  color = texture(textureSampler, UV).rgb;
}`
	result := loadFragmentShader(TEXTURE)

	if result != expected {
		t.Errorf("Fragment shader invalid.\n\n***Expected***:\n%v\n\n***got***:\n%v", expected, result)
	}
}

func TestVertexShaderWithTexture(t *testing.T) {
	expected := `#version 330 core

layout(location = 0) in vec3 vertexPosition_modelspace;
layout(location = 1) in vec2 vertexUV;

out vec2 UV;

uniform mat4 MVP;

void main() {
  gl_Position =  MVP * vec4(vertexPosition_modelspace,1);
  UV = vertexUV;
}`
	result := loadVertexShader(TEXTURE)

	if result != expected {
		t.Errorf("Vertex shader invalid.\n\n***Expected***:\n%v\n\n***got***:\n%v", expected, result)
	}
}

func TestColorShaderCompiles(t *testing.T) {
	if !glfw.Init() {
		t.Errorf("Can't open GLFW")
		return
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	window, err := glfw.CreateWindow(100, 100, "Test", nil, nil)
	if err != nil {
		t.Error(err)
	}
	window.MakeContextCurrent()

	if gl.Init() != 0 {
		t.Errorf("Could not initialise glew.")
	}
	gl.GetError()

	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()

	MakeProgram(COLOR)
}

func TestTextureShaderCompiles(t *testing.T) {
	if !glfw.Init() {
		t.Errorf("Can't open GLFW")
		return
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	window, err := glfw.CreateWindow(100, 100, "Test", nil, nil)
	if err != nil {
		t.Error(err)
	}
	window.MakeContextCurrent()

	if gl.Init() != 0 {
		t.Errorf("Could not initialise glew.")
	}
	gl.GetError()

	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()

	MakeProgram(TEXTURE)
}
