package three

type Vector [4]float32

func (v *Vector) SetX(x float32) {
	v[0] = x
}
func (v *Vector) SetY(y float32) {
	v[1] = y
}
func (v *Vector) SetZ(z float32) {
	v[2] = z
}
