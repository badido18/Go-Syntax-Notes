package syntax

import (
	"fmt"
	"math"
)

func TestVarDeclaration() {
	a := "this i a string"
	var b int = 5
	c := true
	d := 5.48
	const e float64 = 6e48

	fmt.Println("Let's GO !", a, b, c, d+e)
	fmt.Printf("%T\n", e)
	fmt.Println(math.Max(math.Round(0.7), math.Floor(2.7)))
	fmt.Println(math.Pi)
}
