// https://tour.golang.org/concurrency/2

// NOTE: By default, sends and receives block until the other side is ready.
package main

import (
	"fmt"
	"time"
)

func sum(arr []int, c chan int) {
	fmt.Println("start sum")
	s := 0
	for _, v := range arr {
		time.Sleep(time.Duration(v) * time.Millisecond)
		s += v
	}
	fmt.Println("get sum", s)
	c <- s
	fmt.Println("sent sum", s)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)

	go sum(a[:3], c)
	go sum(a[3:], c)

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("s1 wait")
	s1 := <-c
	fmt.Println("s1", s1)

	fmt.Println("s2 wait")
	s2 := <-c
	fmt.Println("s2", s2)

	fmt.Println(s1, s2, s1+s2)

	time.Sleep(1000 * time.Millisecond)

}
