package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	a := math.Pow(1, 2)
	b := math.Pow(2, 3)
	c := math.Pow(3, 4)

	fmt.Println(a, b, c)
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
