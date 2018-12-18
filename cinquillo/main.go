package main

import (
	"fmt"
)

func main() {
	es := estadoInicial([]int{0, 1, -1, 1, 0, 1, 0, 1, 0, 1},
		[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0})

	fmt.Println(es)
}
