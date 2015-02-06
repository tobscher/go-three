# go-three [![Build Status](https://travis-ci.org/tobscher/go-three.svg?branch=master)](https://travis-ci.org/tobscher/go-three) [![GoDoc](https://godoc.org/github.com/tobscher/go-three?status.svg)](https://godoc.org/github.com/tobscher/go-three)

go-three provides a simple API to create and display animated 3D computer graphics.

## Prerequisites

* [GLFW](http://www.glfw.org/) ~> 3.0
* [GLEW](http://glew.sourceforge.net/) >= 1.5.4

NOTE: Check out the `scripts` directory on how to install dependencies.

## Installation

Once you have installed the prerequisites you can install the package via `go get`:
```
go get github.com/Tobscher/go-three
```

## Usage

```go
package main

import (
	"log"

	three "github.com/tobscher/go-three"
	"github.com/tobscher/go-three/geometries"
)

const (
	fov    = 75.0
	width  = 640
	height = 480
	near   = 1
	far    = 10000
)

func main() {
	window, err := three.NewWindow(width, height, "Example")
	if err != nil {
		log.Fatal(err)
	}

	renderer, err := three.NewRenderer(window)
	if err != nil {
		log.Fatal(err)
	}

	scene := three.NewScene()
	camera := three.NewPerspectiveCamera(fov, width/height, near, far)
	camera.Transform.TranslateZ(1000)

	box := geometries.NewCube(200)
	red := three.NewBasicMaterial()
	red.SetColor(three.NewColor(1.0, 0.0, 0.0))

	mesh := three.NewMesh(box, red)

	scene.Add(mesh)

	for !window.ShouldClose() {
		mesh.Transform.RotateX(0.01)
		mesh.Transform.RotateY(0.02)

		renderer.Render(scene, camera)
	}

	renderer.Unload(scene)

	renderer.OpenGLSentinel()
}
```

## Documentation

Documentation can be found on [godoc.org](http://godoc.org/github.com/Tobscher/go-three).

## Examples

Examples can be found in the `examples/` subdirectory.

## Contributing

If you encounter any issues please [file an issue](https://github.com/Tobscher/go-three/issues/new).

Pull requests are welcome.
