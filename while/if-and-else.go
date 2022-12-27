package main

import (
	"fmt"
	"math"
)

//在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。

func Pow(x, y, n float64) float64 {
	if v := math.Pow(x, y); v < n {
		return v
	} else {
		fmt.Printf("%g > %g \n", v, n)
	}
	return n
}

func main() {

	fmt.Println(Pow(3, 2, 10), Pow(3, 3, 10))
}
