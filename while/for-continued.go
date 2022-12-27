package main

import "fmt"

// 初始化语句和后置语句是可选的。
func main() {
	sum := 1
	//此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。
	for sum < 100 {
		sum += sum
	}

	fmt.Println(sum)
}
