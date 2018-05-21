package cmux

import "io"

func (rt *radixTree) matchPrefix(r io.Reader) bool {
	buf := make([]byte, rt.max)
	n, _ := io.ReadFull(r, buf)
	return rt.root.match(buf[:n], true)
}
