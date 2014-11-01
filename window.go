package three

import (
	"errors"
	"fmt"
	glfw "github.com/go-gl/glfw3"
	"runtime"
)

// Window holds information about the dimensions and title of a window.
type Window struct {
	Width  int
	Height int
	Title  string
	window *glfw.Window
}

// NewWindow creates a new window for the given dimensions and title.
// The window is created via GLFW.
func NewWindow(width, height int, title string) (*Window, error) {
	runtime.LockOSThread()

	// Error callback
	glfw.SetErrorCallback(errorCallback)

	// Init glfw
	if !glfw.Init() {
		return nil, errors.New("Could not initialise GLFW.")
	}

	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	// Create window
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, err
	}
	window.SetKeyCallback(keyCallback)
	window.MakeContextCurrent()

	// Use vsync
	glfw.SwapInterval(1)

	w := Window{
		window: window,
		Width:  width,
		Height: height,
		Title:  title,
	}

	return &w, nil
}

// Unload terminates GLFW and closes the current window.
func (w *Window) Unload() {
	glfw.Terminate()
}

// ShouldClose indicates if the OS has received a signal to close this window.
func (w *Window) ShouldClose() bool {
	return w.window.ShouldClose()
}

// Swap swaps buffers and polls events.
func (w *Window) Swap() {
	w.window.SwapBuffers()
	glfw.PollEvents()
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}
