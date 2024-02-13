package main

import (
	"fmt"
	"math"
	_ "strings"

	m "example.local/calc/math"
)

func main() {
	fmt.Println("Hello! 02 first mod on branch 02")
	fmt.Println("Hello! 01 first mod")
	fmt.Println("Hello! 03 first mod")
	slice := []float64{1.0, 2.0, 3.0, 4.0}
	fmt.Println(m.Add(1, 2))
	fmt.Println(m.Sub(1, 2))
	fmt.Println(m.Mul(1, 2))
	fmt.Println(m.Div(1, 2))
	fmt.Println(m.Avg(slice))
	fmt.Println(m.Pi)
	fmt.Println(math.Pi)
	fmt.Println("Hello! 01 master mod")
}
