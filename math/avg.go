package math

func Avg(slice []float64) float64 {
	fmt.Println("Hello! 01 master mod")
	total := 0.0
	for _, slc := range slice {
		total += slc
	}
	return total / float64(len(slice))
}
