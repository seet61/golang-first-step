package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var testsAverage = []testpair{
	{[]float64{}, 0.0},
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

var testsMin = []testpair{
	{[]float64{}, 0.0},
	{[]float64{1, 2}, 1},
	{[]float64{0, 1, 2, 3, 4, 5}, 0},
	{[]float64{-1, 1}, -1},
}

var testsMax = []testpair{
	{[]float64{}, 0.0},
	{[]float64{1, 2}, 2},
	{[]float64{0, 1, 2, 3, 4, 5}, 5},
	{[]float64{-1, 1}, 1},
}

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got", v)
	}
}

func TestAverageTests(t *testing.T) {
	for _, pair := range testsAverage {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	var v float64
	v = Min([]float64{1, 2})
	if v != 1 {
		t.Error("Expected 1, got", v)
	}
}

func TestMinTests(t *testing.T) {
	for _, pair := range testsMin {
		v := Min(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	var v float64
	v = Max([]float64{1, 2})
	if v != 2 {
		t.Error("Expected 2, got", v)
	}
}

func TestMaxTests(t *testing.T) {
	for _, pair := range testsMax {
		v := Max(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
