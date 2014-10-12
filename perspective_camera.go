package three

type PersepectiveCamera struct {
	Position Vector
}

func NewPerspectiveCamera(fov int, aspect, near, far float32) PersepectiveCamera {
	return PersepectiveCamera{}
}
