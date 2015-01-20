package three

import (
	glfw "github.com/go-gl/glfw3"
)

var logger = NewLogger("[go.three] ")

func GetTime() float64 {
	return glfw.GetTime()
}
