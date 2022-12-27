package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	temp := 0.0
	for {
		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
		if math.Abs(z-temp) < 0.00000000001 { //Abs方法 返回 x 的绝对值。
			break
		} else {
			temp = z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
