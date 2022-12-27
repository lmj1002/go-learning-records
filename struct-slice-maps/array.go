package main

import (
	"fmt"
)

// 类型 [n]T 表示拥有 n 个 T 类型的值的数组。
func main() {
	var a [2]string
	a[0] = "test1"
	a[1] = "test2"
	b := [5]int{12, 23, 34, 45, 56}
	fmt.Println(a, b)
}
