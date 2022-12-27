package main

import "fmt"

// 切片文法类似于没有长度的数组文法。
func main() {
	q := []int{2, 3, 5, 7, 11, 13}

	r := []bool{true, false, true, false, false}

	s := []struct {
		i int
		b bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
	}

	fmt.Println(q, r, s)
}
