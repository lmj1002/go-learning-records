package main

import "fmt"

//nil接口值既不保存值也不保存具体类型。
//为nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {

	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {

}
