package runtime

import "fmt"

type ContentAddress struct {
	Line     int
	Col      int
	Filename string
}

func (c ContentAddress) String() string {
	return fmt.Sprintf("(Loc :line %d :col %d :filename %q)", c.Line, c.Col, c.Filename)
}

// used in compile error messages
func (c ContentAddress) Render() string {
	return fmt.Sprintf("file %s @ line %d, col %d", c.Filename, c.Line, c.Col)
}
