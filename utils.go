package three

import "github.com/go-gl/mathgl/mgl32"

func vec2ToFloats(vectors []mgl32.Vec2) []float32 {
	result := []float32{}

	for _, vector := range vectors {
		result = append(result, vector.X(), vector.Y())
	}

	return result
}
