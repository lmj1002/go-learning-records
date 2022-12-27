package main

import (
	"fmt"
	"math"
	"strconv"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(x)
	}
	return fmt.Sprintf(strconv.FormatFloat(math.Sqrt(x), 'g', 0, 32))
}

// Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
func main() {
	fmt.Println(sqrt(2), sqrt(-2))
}
