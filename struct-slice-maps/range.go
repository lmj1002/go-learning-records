package main

import "fmt"

// for 循环的 range 形式可遍历切片或映射。
func main() {
	//当使用 for 循环遍历切片时，每次迭代都会返回两个值。
	//第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
	var pow = []int{1, 2, 3, 4, 5, 6, 7, 8}
	for k, v := range pow {
		fmt.Printf("2**%d = %d\n", k, v)
	}
	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, v := range pow { //使用_忽略掉不需要的值
		fmt.Printf("%d\n", v)
	}
}
