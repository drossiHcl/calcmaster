package math

func Avg(slice []float64) float64 {
	total := 0.0
	for _, slc := range slice {
		total += slc
	}
	return total / float64(len(slice))
}
