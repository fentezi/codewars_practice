package main

import (
	"fmt"
	"math"
)

func EquableTriangle(a, b, c int) bool {
	s := float64(a+b+c) / 2.0
	area := math.Sqrt(s * (s - float64(a)) * (s - float64(b)) * (s - float64(c)))
	perimeter := float64(a + b + c)
	return math.Round(area) == math.Round(perimeter)
}

func main() {
	fmt.Println(EquableTriangle(5, 12, 13))
}
