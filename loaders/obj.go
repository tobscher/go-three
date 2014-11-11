package loaders

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/tobscher/go-three"
)

// LoadFromObj loads an obj file and returns a Geometry.
func LoadFromObj(path string) (*three.Geometry, error) {
	// Load file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("(.*?) (.*)")

	vertices := make([]mgl32.Vec3, 0)
	faces := make([]*three.Face, 0)

	// Scan lines
	for fileScanner.Scan() {
		text := fileScanner.Text()

		// Match header and rest of line
		result := r.FindStringSubmatch(text)
		header := result[1]
		restOfLine := result[2]

		// Handle each line indivdual
		switch header {
		case "v":
			// Vertex line
			vert := mgl32.Vec3{}
			count, _ := fmt.Sscanf(restOfLine, "%f %f %f", &vert[0], &vert[1], &vert[2])
			if count != 3 {
				return nil, errors.New("Invalid obj file. Vertex line should be of format 'x y z'")
			}
			vertices = append(vertices, vert)
		case "f":
			f := make([]uint16, 3)
			count, _ := fmt.Sscanf(restOfLine, "%d %d %d", &f[0], &f[1], &f[2])
			if count != 3 {
				return nil, errors.New("Invalid obj file. Face line should be of format 'a b c'")
			}

			faces = append(faces, three.NewFace(f[0]-1, f[1]-1, f[2]-1))
		default:
			// eat line
		}
	}

	obj := &three.Geometry{}
	obj.SetVertices(vertices)
	obj.SetFaces(faces)

	return obj, nil
}
