package main

import "fmt"

// 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。
// 在实践中，切片比数组更常用。
func main() {
	primes := [6]int{1, 2, 3, 4, 5, 6}
	//它会选择一个半开区间，包括第一个元素，但排除最后一个元素。
	//以下表达式创建了一个切片，它包含 primes中下标从 1 到 3 的元素
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
