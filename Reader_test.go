package bufio

import (
	"bytes"
	"fmt"
	"io"
)

func ExampleReader() {
	var r = NewReader(bytes.NewBufferString("Hello, World!"))
	var p, _ = r.Peek(5)
	fmt.Println(string(p))
	p = make([]byte, 20)
	var n, err = io.ReadFull(r, p)
	fmt.Println(n)
	fmt.Println(err == io.ErrUnexpectedEOF)
	fmt.Println(string(p[:5]))
	// Output:
	// Hello
	// 13
	// true
	// Hello
}
