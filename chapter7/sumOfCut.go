package main

import "fmt"

func sum(xs []int) int {
	i := int(0)
	for _, val := range xs {
		i += val
	}
	return i
}

func main() {
	x := []int{1, 2, 4, 6, 8}
	fmt.Println(sum(x))
}
