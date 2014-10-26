package three

import (
	"os"
	"testing"
)

func TestColorShaderCompiles(t *testing.T) {
	if os.Getenv("SKIP_GLFW") != "" {
		t.Skip()
	}

	RunIn3DContext(func() {
		MakeProgram(COLOR)
	})
}

func TestTextureShaderCompiles(t *testing.T) {
	if os.Getenv("SKIP_GLFW") != "" {
		t.Skip()
	}

	RunIn3DContext(func() {
		MakeProgram(TEXTURE)
	})
}
