package main

import "fmt"

func half(x uint) (uint, bool) {
	i := x % 2
	if i == 0 {
		return 0, true
	}
	return 1, false
}

func main() {
	fmt.Print("half(1): ")
	fmt.Println(half(1))
	fmt.Print("half(2): ")
	fmt.Println(half(2))
	fmt.Print("half(3): ")
	fmt.Println(half(3))
}
