package statistic

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Sum(arr []float64) float64 {
	res := float64(0)
	for _, val := range arr {
		res += val
	}
	return res
}
