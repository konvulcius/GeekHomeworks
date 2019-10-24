package quadrequation

import (
	"testing"
)

type result struct {
	solve [2]float64
	err   error
}

type testvalue struct {
	values []float64
	result
}

var test = []testvalue{
	{[]float64{1, -5, 4}, result{[2]float64{4, 1}, nil}},
	{[]float64{3, -14, -5}, result{[2]float64{5, -0.3333333333333333}, nil}},
	{[]float64{0, 4, -2}, result{[2]float64{0.5, 0.5}, nil}},
}

func TestQuadr(t *testing.T) {
	for _, temp := range test {
		x1, x2, err := quadrEquation(temp.values[0], temp.values[1], temp.values[2])
		if x1 != temp.result.solve[0] || x2 != temp.result.solve[1] {
			t.Error(
				"need", temp.result.solve[0], temp.result.solve[1],
				"yours", x1, "and", x2,
			)
		} else if err != temp.result.err {
			t.Error(
				"dont be", err,
			)
		}
	}
}
