package main

import "fmt"

// Vertex 一个结构体（struct）就是一组字段（field）。
type Vertex struct {
	x int
	Y int
}

func main() {
	//结构体文法通过直接列出字段的值来新分配一个结构体。
	v := Vertex{1, 2}
	x := Vertex{x: 1} // X:0 Y:1
	y := Vertex{}     // X:0 Y:0
	v.x = 4           //结构体字段使用点号来访问。
	fmt.Println(v.x, v.Y)
	//特殊的前缀 & 返回一个指向结构体的指针。
	n := &Vertex{1, 2}

	//结构体字段可以通过结构体指针来访问。
	p := &v
	p.Y = 10086
	//如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
	fmt.Println(p, v, x, y, n)
}
