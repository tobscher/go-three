package loaders

type StringReader string

func (s StringReader) Read(byt []byte) (n int, err error) {
	copy(byt, string(s))
	return len(byt), nil
}
