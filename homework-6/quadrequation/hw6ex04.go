package quadrequation

import (
	"fmt"
	"math"
)

type myError struct {
	Some string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%v", e.Some)
}

func wrong(s string) error {
	return &myError{
		s,
	}
}

func quadrEquation(a, b, c float64) (float64, float64, error) {
	var x1, x2 float64

	if a == 0 && b == 0 {
		return 0, 0, wrong("It is not quadratic equation")
	}

	d := b*b - 4*a*c
	if d < 0 {
		return 0, 0, wrong("There is no answers")
	} else if d == 0 {
		x1 = (-b) / (2 * a)
		return x1, x1, nil
	} else {
		if a != 0 {
			x1 = ((-b) + math.Sqrt(d)) / (2 * a)
			x2 = ((-b) - math.Sqrt(d)) / (2 * a)
		} else {
			x1 = (-c) / b
			x2 = x1
		}
		return x1, x2, nil
	}
}
