package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) Bounds() image.Rectangle {
	//TODO implement me
	panic("implement me")
}

func (i Image) At(x, y int) color.Color {
	//TODO implement me
	panic("implement me")
}

func (i Image) ColorModel() color.Model {
	//TODO implement me
	panic("implement me")
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
