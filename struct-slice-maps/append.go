package main

// 为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数。
// 内建函数的https://go-zh.org/pkg/builtin/#append对此函数有详细的介绍。
func main() {
	var s []int //添加一个空切片
	PrintSlice(s)
	s = append(s, 1) //这个切片会按需增长
	PrintSlice(s)
	s = append(s, 2, 3, 4) //一次性添加多个元素
	PrintSlice(s)
}
