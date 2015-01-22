package three

import (
	"errors"
	"fmt"

	glfw "github.com/go-gl/glfw3"
)

var currentWindow *Window

// Window holds information about the dimensions and title of a window.
type Window struct {
	Settings WindowSettings
	window   *glfw.Window
}

// WindowSettings holds information that describe how the window should constructed.
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
	logger.Debug("Initializing GLFW")
	if !glfw.Init() {
		return nil, errors.New("Could not initialise GLFW.")
	}

	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)
	glfw.WindowHint(glfw.OpenglDebugContext, 1)

	var monitor *glfw.Monitor
	var err error
	if settings.Fullscreen {
		logger.Debug("Get primary monitor to create fullscreen window.")
		monitor, err = glfw.GetPrimaryMonitor()
		if err != nil {
			return nil, err
		}

		logger.Debug("Checking available video modes:")
		videoModes, err := monitor.GetVideoModes()
		if err != nil {
			return nil, err
		}

		for _, videoMode := range videoModes {
			logger.Debug(fmt.Sprintf("-- %+v", videoMode))
		}

		idealVideoMode := videoModes[len(videoModes)-1]

		settings.Width = idealVideoMode.Width
		settings.Height = idealVideoMode.Height
	}

	// Create window
	logger.Info("Creating new window")
	window, err := glfw.CreateWindow(settings.Width, settings.Height, settings.Title, monitor, nil)
	if err != nil {
		return nil, err
	}
	window.SetKeyCallback(keyCallback)
	window.MakeContextCurrent()
	window.SetInputMode(glfw.StickyKeys, 1)

	// Use vsync
	glfw.SwapInterval(1)

	w := &Window{
		window:   window,
		Settings: settings,
	}

	currentWindow = w

	return w, nil
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

// Width returns the width of the window in pixels.
func (w *Window) Width() int {
	width, _ := w.window.GetSize()
	return width
}

// Height returns the height of the window in pixels.
func (w *Window) Height() int {
	_, height := w.window.GetSize()
	return height
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}
