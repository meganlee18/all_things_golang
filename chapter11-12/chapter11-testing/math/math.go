package math

/* Writing a good suite of tests is not always easy,
but the process of writings tests often reveals
more about a problem then you may at first realize.
For example, with our Average function
what happens if you pass in an empty list
([]float64{})? How could we modify the function
to return 0 in this case? */

func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

/* Do a loop through the slice to check if it's larger than first value. Throw it out until the last number is left */
func Max(numbers []float64) float64 {
	var number float64
	for _, value := range numbers {
		if value >= numbers[0] {
			number = value
		}
	}
	return number
}

func Min(numbers []float64) float64 {
	var number float64
	for _, value := range numbers {
		if value <= numbers[0] {
			number = value
		}
	}
	return number
}
