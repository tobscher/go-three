package three

// Geometry is the interface which defines the shape of a
// 3D object.
//
// generateVertexBuffer returns an array of points which define the shape
// of the underlying geometry.
type Geometry interface {
	generateVertexBuffer() []float32
}
