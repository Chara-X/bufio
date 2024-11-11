package bufio

import (
	"bufio"
	"io"
)

type Writer struct {
	b   *bufio.Writer
	w   io.Writer
	buf []byte
	n   int
}

func NewWriter(w io.Writer) *Writer {
	if Reference {
		return &Writer{b: bufio.NewWriter(w)}
	}
	return &Writer{w: w, buf: make([]byte, 4096)}
}
func (b *Writer) Write(p []byte) (n int, err error) {
	if Reference {
		return b.b.Write(p)
	}
	n = len(p)
	for len(p) > 0 {
		var n = copy(b.buf[b.n:], p)
		b.n += n
		p = p[n:]
		if b.n == len(b.buf) {
			b.Flush()
		}
	}
	return n, nil
}
func (b *Writer) Flush() error {
	if Reference {
		return b.b.Flush()
	}
	b.w.Write(b.buf[:b.n])
	b.n = 0
	return nil
}
