package main

import (
	"fmt"
	"math"
)

func main() {
	var Grade []string
	Grade = []string{"F", "D", "C", "B", "A"}
	fmt.Println(Grade)
	score := 80
	sl := math.Ceil(float64(score)/20) - 1
	fmt.Println(sl)
	// s := 4
	fmt.Println(Grade[int(sl)])
	// fmt.Println(math.Ceil(float64(score) / 10))
	// var buffer int
	// fmt.Println("buffer",score(10))
	// fmt.Println("Hi")
}
