package loaders

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

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
	normals := make([]mgl32.Vec3, 0)
	faces := make([]*three.Face, 0)

	// Scan lines
	for fileScanner.Scan() {
		text := fileScanner.Text()

		// Match header and rest of line
		result := r.FindStringSubmatch(text)

		if len(result) != 3 {
			log.Println("Skip line. Wrong format.")
			continue
		}

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
		case "vn":
			// Normal line
			normal := mgl32.Vec3{}
			count, _ := fmt.Sscanf(restOfLine, "%f %f %f", &normal[0], &normal[1], &normal[2])
			if count != 3 {
				return nil, errors.New("Invalid obj file. Normal line should be of format 'x y z'")
			}
			normals = append(normals, normal)
		case "f":
			f := []uint16{}
			n := []uint16{}

			faceElements := strings.Split(restOfLine, " ")
			if len(faceElements) < 3 {
				return nil, errors.New("Invalid obj file. Face line should be of format 'a b c [d]'")
			}

			for _, element := range faceElements {
				elementTypes := strings.Split(element, "/")
				if len(elementTypes) < 1 {
					return nil, errors.New("Invalid obj file. Face element has wrong format 'v[[/vn][/vt]]'")
				}

				i, err := strconv.Atoi(elementTypes[0])
				if err != nil {
					return nil, errors.New("Invalid obj file. Face vertex index is not an integer.")
				}
				f = append(f, uint16(i)-1)

				// Process normal index
				if len(elementTypes) > 2 {
					i, err = strconv.Atoi(elementTypes[2])
					if err != nil {
						return nil, errors.New("Invalid obj file. Face normal index is not an integer.")
					}
					n = append(n, uint16(i)-1)
				}
			}

			for i := 1; i < len(f)-1; i++ {
				face := three.NewFace(f[0], f[i], f[i+1])

				if len(n) > 0 {
					face.AddNormal(n[0], n[i], n[i+1])
				}

				faces = append(faces, face)
			}
		default:
			// eat line
		}
	}

	obj := &three.Geometry{}
	obj.SetVertices(vertices)
	obj.SetNormals(normals)
	obj.SetFaces(faces)

	log.Println("Obj loading report:")
	log.Printf("-- Vertices: %v\n", len(obj.Vertices()))
	log.Printf("-- UVs: %v\n", len(obj.UVs()))
	log.Printf("-- Normals: %v\n", len(obj.Normals()))
	log.Printf("-- Faces: %v\n", len(obj.Faces()))

	return obj, nil
}
