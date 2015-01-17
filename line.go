package three

import "github.com/go-gl/gl"

// Line is a representation of a line in 3D space. It consists
// of a geometry which defines all vertices of the lines to render
// and a material. Lines can be transformed.
type Line struct {
	Object3D

	geometry Shape
	material Appearance
}

// NewLine creates a new line for the given geometry (set of vertices) and material.
func NewLine(geometry Shape, material Appearance) *Line {
	l := Line{
		Object3D: NewObject3D(),
		geometry: geometry,
		material: material,
	}

	l.vertexBuffer = newVertexBuffer(geometry)

	return &l
}

// Geometry returns an object which describes the shape of the line.
func (l *Line) Geometry() Shape {
	return l.geometry
}

// Material returns an objects which described the appaerance of the line.
func (l *Line) Material() Appearance {
	return l.material
}

// Mode returns an enum which is used to define the type of primitive to render.
func (l *Line) Mode() gl.GLenum {
	return gl.LINES
}
