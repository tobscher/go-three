package three

type Color [3]float32

func (c Color) R() float32 {
	return c[0]
}

func (c Color) G() float32 {
	return c[1]
}

func (c Color) B() float32 {
	return c[2]
}
