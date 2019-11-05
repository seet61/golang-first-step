package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	var y = make([]int, 3, 9)
	fmt.Println(y)

	z := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(z[2:5])

}
