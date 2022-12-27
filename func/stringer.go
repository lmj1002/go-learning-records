package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v(%v years)", p.Name, p.Age)
}

//  fmt 包中定义的 Stringer 是最普遍的接口之一。
// Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。

func main() {
	a := Person{"test", 18}
	z := Person{"liudehua", 9001}

	fmt.Println(a, z)
}
