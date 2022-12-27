package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap = %d %v \n", len(s), cap(s), s)
}

// 在进行切片时，你可以利用它的默认行为来忽略上下界。
// 切片下界的默认值为 0，上界则是该切片的长度。
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	//a := s[:5]
	//b := s[2:]
	//c := s[:]
	//fmt.Println(a, b, c)
	s = s[:0] // 截取切片使其长度为 0
	printSlice(s)
	s = s[:4] // 拓展其长度
	printSlice(s)
	s = s[2:] // 舍弃前两个值
	printSlice(s)
}
