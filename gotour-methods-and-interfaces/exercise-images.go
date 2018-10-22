package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect{0, 0, 100, 100}
}

func (i Image) At(x, y int) color.Color {
	if x == y {
		return color.RGBA{1, 1, 255, 255}
	}

	return color.RGBA{255, 255, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
