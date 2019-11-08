package math

import "rsc.io/quote"

func Hello() string {
	return quote.Hello()
}

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//Поиск наименьшего числа в массиве
func Min(xs []float64) float64 {
	min := 99999999.0
	for _, val := range xs {
		if val < min {
			min = val
		}
	}
	return min
}

//Поиск наибольшего числа в массиве
func Max(xs []float64) float64 {
	max := 0.0
	for _, val := range xs {
		if val > max {
			max = val
		}
	}
	return max
}
