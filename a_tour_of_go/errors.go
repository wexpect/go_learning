// https://tour.golang.org/methods/20

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return x, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
