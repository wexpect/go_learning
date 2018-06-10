// https://tour.golang.org/methods/22

package main

import "fmt"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, err error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	r := MyReader{}
	b := make([]byte, 8)
	n, err := r.Read(b)
	fmt.Printf("%v %v %v %q\n", n, err, b, b)
}
