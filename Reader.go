package bufio

import (
	"bufio"
	"io"
)

type Reader struct {
	b    *bufio.Reader
	r    io.Reader
	buf  []byte
	i, j int
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
	if b.i == b.j {
		b.i = 0
		b.j, _ = b.r.Read(b.buf)
		if b.j == 0 {
			return 0, io.EOF
		}
	}
	n = copy(p, b.buf[b.i:b.j])
	b.i += n
	return n, nil
}
func (b *Reader) Peek(n int) ([]byte, error) {
	if Reference {
		return b.b.Peek(n)
	}
	for b.j-b.i < min(n, len(b.buf)) {
		copy(b.buf, b.buf[b.i:b.j])
		b.i = 0
		b.j -= b.i
		var n, _ = b.r.Read(b.buf[b.j:])
		if n == 0 {
			return b.buf[:b.j], io.ErrUnexpectedEOF
		}
		b.j += n
	}
	return b.buf[b.i : b.i+n], nil
}
