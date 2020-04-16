package syntax

import "fmt"

//closures

func Closure() func(string) int {
	Nmes := ""
	return func(a string) int {
		Nmes += a
		return len(Nmes)
	}
}

func TestClosure() {
	na := Closure()
	fmt.Println(na("Madjid"))

}
