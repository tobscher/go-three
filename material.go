package three

// Appearance is the interface which defines the appearance of 3D objects.
// For example: solid color, textured, etc.
//
// SetProgram is used to store the given program and to cache it inside
// the material.
// Program returns a program which is used to define the appearance of
// the 3D object. A program consists of a vertex and a fragment shader.
type Appearance interface {
	SetProgram(*Program)
	Program() *Program
}

// Colored is an interface that indicates that a material can have a solid color.
type Colored interface {
	Color() *Color
}

// Textured is an interface that indicates that a material can have a texture.
type Textured interface {
	Texture() *Texture
	SetTexture(*Texture)
}

// Wireframed is an interface that indicates that a material can be rendered
// as wireframes only.
type Wireframed interface {
	Wireframe() bool
}

// Texter is an interface composition of Appaerance and Textured.
type Texter interface {
	Appearance
	Textured
}
