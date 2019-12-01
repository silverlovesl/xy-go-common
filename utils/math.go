package utils

import "math"

// Round はplaces の位置で四捨五入した値を返却
func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}

// SumFloat リストの合計値を計算
func SumFloat(values []float64) float64 {

	var sum float64

	for _, val := range values {
		sum += val
	}
	return sum
}
