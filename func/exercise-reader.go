package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (m MyReader) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法

func main() {
	reader.Validate(MyReader{})
}
