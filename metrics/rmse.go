package metrics

import (
	"math"
)

func RMSE(original, predicted []float64) float64 {
	var sum float64
	if len(original) != len(predicted) {
		panic("RMSE - parameter length mismatch")
	}

	for i := range original {
		predError := predicted[i] - original[i]
		sum += math.Pow(predError, 2)
	}
	meanError := sum / float64(len(original))
	return math.Sqrt(meanError)
}
