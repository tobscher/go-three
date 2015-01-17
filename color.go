package three

// Color stores color information for the red, green and blue channel.
type Color [3]float32

// R returns the value for the red channel.
// Value can be between 0 and 1
func (c Color) R() float32 {
	return c[0]
}

// G returns the value for the green channel.
// Value can be between 0 and 1
func (c Color) G() float32 {
	return c[1]
}

// B returns the value for the blue channel.
// Value can be between 0 and 1
func (c Color) B() float32 {
	return c[2]
}

// Float returns the current as array of float32's
func (c Color) Float() []float32 {
	return []float32{
		c[0],
		c[1],
		c[2],
	}
}
