package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	a := rand.Intn(10)
	b := rand.Intn(10)
	sum := add(a, b)

	fmt.Print("My favorite number is ", sum)
}
