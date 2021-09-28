package main

import (
	"fmt"
	"math"
	"math/rand"
)

func cubic(x float64) (y float64) {
	y = math.Pow(x, 3)
	return
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Three random numbers are", rand.Intn(100), rand.Intn(100), rand.Intn(100))
	fmt.Println("Pi is", math.Pi)
	println(cubic(3))
}
