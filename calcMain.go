package main

import (
	"fmt"
	"math"
	_ "strings"

	m "example.local/calc/math"
)

func main() {
	fmt.Println("Hello!")
	slice := []float64{1.0, 2.0, 3.0, 4.0}
	fmt.Println(m.Add(1, 2))
	fmt.Println(m.Sub(1, 2))
	fmt.Println(m.Mul(1, 2))
	fmt.Println(m.Div(1, 2))
	fmt.Println(m.Avg(slice))
	fmt.Println(m.Pi)
	fmt.Println(math.Pi)
}
