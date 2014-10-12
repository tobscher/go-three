package three

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type VertexCollection struct {
	items []mgl32.Vec3
}

func (vc *VertexCollection) Push(vector mgl32.Vec3) {
	vc.items = append(vc.items, vector)
}

type BoxGeometry struct {
	program          gl.Program
	vertexBufferData [36]float32
	Vertices         VertexCollection
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