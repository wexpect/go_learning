// https://tour.golang.org/concurrency/8

package main

import (
	"fmt"
	"reflect"

	"golang.org/x/tour/tree"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		fmt.Printf("left of %v start \n", t.Value)
		Walk(t.Left, ch)
		fmt.Printf("left of %v finish \n", t.Value)
	}

	fmt.Printf("node %v start \n", t.Value)
	ch <- t.Value
	fmt.Printf("node %v finish \n", t.Value)

	if t.Right != nil {
		fmt.Printf("right of %v start \n", t.Value)
		Walk(t.Right, ch)
		fmt.Printf("right of %v finish \n", t.Value)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int)
	vs1 := make([]int, 0, 10)
	go Walk(t1, ch1)

	ch2 := make(chan int)
	vs2 := make([]int, 0, 10)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		v1 := <-ch1
		vs1 = append(vs1, v1)

		v2 := <-ch2
		vs2 = append(vs2, v2)
	}

	fmt.Println("vs1", vs1)
	fmt.Println("vs2", vs2)
	return reflect.DeepEqual(vs1, vs2)

}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(3)))
}
