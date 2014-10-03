package main

import (
	"fmt"
)

// START1 OMIT
func f(j int, g func(int) int) int {
	return g(j)
}

//END1 OMIT
//START2 OMIT
func main() {
	var t = func(j int) int {
		return j + 1
	}

	var g = func(j int) int {
		return j * j
	}
	fmt.Println(f(f(1, t), g)) // HL
}

//END2 OMIT
