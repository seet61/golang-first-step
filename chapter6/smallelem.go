package main

import "fmt"

func main() {
	var y = 999
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	for _, value := range x {
		if value < y {
			y = value
		}
	}
	fmt.Println("y:", y)
}
