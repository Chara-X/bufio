package bufio

import (
	"bytes"
	"fmt"
)

func ExampleWriter() {
	var w = bytes.NewBuffer(nil)
	var b = NewWriter(w)
	b.Write([]byte("Hello, World!"))
	b.Flush()
	b.Write([]byte("Hello, World!"))
	fmt.Println(w.String())
	// Output:
	// Hello, World!
}
