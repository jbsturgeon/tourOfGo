package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(b []byte) (int, error) {
	bRead, err := rr.r.Read(b)
	for ii := 0; ii < len(b); ii++ {
		if (b[ii] >= 'A' && b[ii] < 'N') || (b[ii] >= 'a' && b[ii] < 'n') {
			b[ii] += 13
		} else if (b[ii] > 'M' && b[ii] <= 'Z') || (b[ii] > 'm' && b[ii] <= 'z') {
			b[ii] -= 13
		}
	}

	return bRead, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	println("\n")
}
