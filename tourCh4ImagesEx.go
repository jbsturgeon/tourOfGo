package main

import (
  "fmt"
  "image"
  "image/color"
  "golang.org/s/tour/pic"
)

type Image struct {
  O_x int
  O_y int
  Wid int
  Hgt int

  clr [][]color.Color
}

func (i *Image) ColorModel() color.Model {
  return  color.RGBAModel
}

func (i *Image) Bounds() Rectangle {
  return image.Rectangle(i.O_x, i.O_y, i.O_x+Wid, i.O_y+Hgt)
}

func (i *Image) At(x,y int) color.Color {

}


func main() {
  m := Image {10, 10, 100, 100, a lot of colors???}
  pic.ShowImage(m)
}
