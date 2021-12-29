package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13Char(b byte) byte {
	var beg byte
	if b >= 'A' && b <= 'Z' {
		beg = 'A'
	} else if b >= 'a' && b <= 'z' {
		beg = 'a'
	} else {
		return b
	}
	return (((b - beg) + 13) % 26) + beg
}
func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := range b {
		b[i] = rot13Char(b[i])
	}
	return n, err;
}
func main() {
	s := strings.NewReader(
		"Yrgf ubcr gung gur Ohaqrfjrue Plore Qrsrafr pragre vf abg hfvat guvf sbe rapelcgvbaf")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)


}
