package three

type VertexCollection struct {
	items []Vector
}

func (vc *VertexCollection) Push(vector Vector) {
	vc.items = append(vc.items, vector)
}

type BoxGeometry struct {
	Vertices VertexCollection
}

func NewBoxGeometry(width, height, depth float32) BoxGeometry {
	geometry := BoxGeometry{}

	// // Triangle 1
	// // -1.0f,-1.0f,-1.0f
	// // -1.0f,-1.0f, 1.0f
	// // -1.0f, 1.0f, 1.0f
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: -1.0, Z: -1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: -1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: 1.0, Z: 1.0})

	// // Triangle 2
	// //  1.0f, 1.0f,-1.0f
	// // -1.0f,-1.0f,-1.0f
	// // -1.0f, 1.0f,-1.0f
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: 1.0, Z: -1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: -1.0, Z: -1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: 1.0, Z: -1.0})

	// // Triangle 3
	// //  1.0f,-1.0f, 1.0f
	// // -1.0f,-1.0f,-1.0f
	// //  1.0f,-1.0f,-1.0f
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: -1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: -1.0, Z: -1.0})
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: -1.0, Z: -1.0})

	// // Triangle 4
	// //  1.0f, 1.0f,-1.0f
	// //  1.0f,-1.0f,-1.0f
	// // -1.0f,-1.0f,-1.0f
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: 1.0, Z: -1.0})
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: -1.0, Z: -1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: -1.0, Z: -1.0})

	// // Triangle 5
	// // -1.0f,-1.0f,-1.0f
	// // -1.0f, 1.0f, 1.0f
	// // -1.0f, 1.0f,-1.0f
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: -1.0, Z: -1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: 1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: 1.0, Z: -1.0})

	// // Triangle 6
	// //  1.0f,-1.0f, 1.0f
	// // -1.0f,-1.0f, 1.0f
	// // -1.0f,-1.0f,-1.0f
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: -1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: -1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: -1.0, Z: -1.0})

	// // Triangle 7
	// // -1.0f, 1.0f, 1.0f
	// // -1.0f,-1.0f, 1.0f
	// //  1.0f,-1.0f, 1.0f
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: 1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: -1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: -1.0, Z: 1.0})

	// // Triangle 8
	// // 1.0f, 1.0f, 1.0f
	// // 1.0f,-1.0f,-1.0f
	// // 1.0f, 1.0f,-1.0f
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: 1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: -1.0, Z: -1.0})
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: 1.0, Z: -1.0})

	// // Triangle 9
	// // 1.0f,-1.0f,-1.0f
	// // 1.0f, 1.0f, 1.0f
	// // 1.0f,-1.0f, 1.0f
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: -1.0, Z: -1.0})
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: 1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: -1.0, Z: 1.0})

	// // Triangle 10
	// //  1.0f, 1.0f, 1.0f,
	// //  1.0f, 1.0f,-1.0f,
	// //  -1.0f, 1.0f,-1.0f,
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: 1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: 1.0, Z: -1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: 1.0, Z: -1.0})

	// // Triangle 11
	// //  1.0f, 1.0f, 1.0f,
	// //  -1.0f, 1.0f,-1.0f,
	// //  -1.0f, 1.0f, 1.0f,
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: 1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: 1.0, Z: -1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: 1.0, Z: 1.0})

	// // Triangle 12
	// //  1.0f, 1.0f, 1.0f,
	// //  -1.0f, 1.0f, 1.0f,
	// //  1.0f,-1.0f, 1.0f
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: 1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: -1.0, Y: 1.0, Z: 1.0})
	// geometry.Vertices.Push(Vector3{X: 1.0, Y: -1.0, Z: 1.0})

	return geometry
}
