package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // & 操作符会生成一个指向其操作数的指针。
	fmt.Println(*p)
	*p = 21 //* 操作符表示指针指向的底层值。
	fmt.Println(i)
	fmt.Println(p)
	p = &j
	*p = *p / 37
	fmt.Println(j)
	fmt.Println(&j, p)
}
