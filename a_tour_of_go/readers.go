// https://tour.golang.org/methods/21

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	r := strings.NewReader("this is test")
	b := make([]byte, 8)

	for true {
		n, err := r.Read(b)
		fmt.Printf("%v %v \n", n, err)
		fmt.Printf("b[:n] %v \n", b[:n])
		fmt.Printf("b[:n] %q \n", b[:n])

		if err == io.EOF {
			break
		}
	}

}
