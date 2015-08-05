package main

import "fmt"

func split(sum int) (x, y int) {
	var a, b int
	a = 3389
	b = a / sum

	x = sum * 4 / 9
	y = sum - x
	return a, b
}

func main() {
	fmt.Println(split(37))
}
