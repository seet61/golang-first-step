package main

import "fmt"

func swap(x1 *int, y1 *int) {
	tmp := *x1
	*x1 = *y1
	*y1 = tmp
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}
