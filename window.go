package three

import (
	"errors"
	"fmt"

	glfw "github.com/go-gl/glfw3"
)

// Window holds information about the dimensions and title of a window.
type Window struct {
	settings WindowSettings
	window   *glfw.Window
}

type WindowSettings struct {
	Width      int
	Height     int
	Title      string
	Fullscreen bool
	ClearColor *Color
}

// NewWindow creates a new window for the given dimensions and title.
// The window is created via GLFW.
func NewWindow(settings WindowSettings) (*Window, error) {
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

	var monitor *glfw.Monitor
	var err error
	if settings.Fullscreen {
		logger.Println("Get primary monitor to create fullscreen window.")
		monitor, err = glfw.GetPrimaryMonitor()
		if err != nil {
			return nil, err
		}

		logger.Println("Checking available video modes:")
		videoModes, err := monitor.GetVideoModes()
		if err != nil {
			return nil, err
		}

		for _, videoMode := range videoModes {
			logger.Printf("-- %+v\n", videoMode)
		}

		idealVideoMode := videoModes[len(videoModes)-1]

		settings.Width = idealVideoMode.Width
		settings.Height = idealVideoMode.Height
	}

	// Create window
	window, err := glfw.CreateWindow(settings.Width, settings.Height, settings.Title, monitor, nil)
	if err != nil {
		return nil, err
	}
	window.SetKeyCallback(keyCallback)
	window.MakeContextCurrent()
	window.SetInputMode(glfw.StickyKeys, 1)

	// Use vsync
	glfw.SwapInterval(1)

	w := Window{
		window:   window,
		settings: settings,
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

// SetTitle sets the window title
func (w *Window) SetTitle(title string) {
	w.window.SetTitle(title)
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}
