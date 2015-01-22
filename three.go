package three

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/tobscher/go-three/logging"
)

var logger = logging.GetLogger("go-three")

// GetTime returns the number of seconds since the timer was started.
//
// Please refer to http://www.glfw.org/docs/latest/input.html#time for more
// information.
func GetTime() float64 {
	return glfw.GetTime()
}
