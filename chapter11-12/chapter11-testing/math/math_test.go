package math

import "testing"

type testpair struct {
	values  []float64
	average float64
	max     float64
	min     float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5, 2, 1},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 1, 1},
	{[]float64{-1, 1}, 0, 1, -1},
}

func TestOneAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5", "instead got ", v)
	}
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
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

/* Write a series of tests for the Min and Max functions
you wrote in the previous chapter */

//individual test
func TestOneMax(t *testing.T) {
	var max float64
	max = Max([]float64{45, 54})
	if max != 54 {
		t.Error("Expected 54", "instead got", max)
	}
}

//individual test
func TestOneMin(t *testing.T) {
	var min float64
	min = Min([]float64{45, 54})
	if min != 45 {
		t.Error("Expected 45", "instead got", min)
	}
}

//testing testpairs above
func TestMax(t *testing.T) {
	for _, pair := range tests {
		v := Max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"instead got", v,
			)
		}
	}
}

//testing testpairs above
func TestMin(t *testing.T) {
	for _, pair := range tests {
		v := Min(pair.values)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"Expected", pair.min,
				"Instead got", v,
			)
		}
	}
}
