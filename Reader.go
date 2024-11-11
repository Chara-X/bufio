package bufio

import (
	"bufio"
	"io"
)

type Reader struct {
	b    *bufio.Reader
	r    io.Reader
	buf  []byte
	i, n int
}

func NewReader(r io.Reader) *Reader {
	if Reference {
		return &Reader{b: bufio.NewReader(r)}
	}
	return &Reader{r: r, buf: make([]byte, 4096)}
}
func (b *Reader) Read(p []byte) (n int, err error) {
	if Reference {
		return b.b.Read(p)
	}
	n = copy(p, b.buf[b.i:b.n])
	b.i += n
	if b.i == b.n {
		b.i = 0
		b.n, _ = b.r.Read(b.buf)
		if b.n == 0 {
			return n, io.EOF
		}
	}
	return n, nil
}
func (b *Reader) Peek(n int) ([]byte, error) {
	if Reference {
		return b.b.Peek(n)
	}
	for b.n-b.i < min(n, len(b.buf)) {
		if b.i > 0 {
			copy(b.buf, b.buf[b.i:b.n])
			b.i = 0
			b.n -= b.i
		}
		var n, _ = b.r.Read(b.buf[b.n:])
		b.n += n
	}
	return b.buf[b.i : b.i+n], nil
}
