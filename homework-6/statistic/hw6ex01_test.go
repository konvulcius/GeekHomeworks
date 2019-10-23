package statistic

import "testing"

type testpair struct {
	values []float64
	result float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

var testSum = []testpair{
	{[]float64{5, 10, 2}, 17},
	{[]float64{100, 200, 45}, 345},
	{[]float64{18, 20, 13}, 51},
}

func TestSum(t *testing.T) {
	for _, temp := range testSum {
		res := Sum(temp.values)
		if res != temp.result {
			t.Error(
				"For", temp.values,
				"expected", temp.result,
				"got", res,
			)
		}
	}
}
