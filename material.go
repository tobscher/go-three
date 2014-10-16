package three

type Material interface {
	Program() Program
	Wireframe() bool
	Attributes() []Attribute
	BufferDataFor(ProgramFeature, Geometry) []float32
}
