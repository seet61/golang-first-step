package main

import "fmt"

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextEven := makeOddGenerator()
	for i := 0; i < 5; i++ {
		fmt.Println(nextEven())
	}
}
