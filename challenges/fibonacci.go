package challenges

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	u0 := 0
	u1 := 1
	return func(e int) int {
		if e == 0 {
			return u0
		} else {
			if e == 1 {
				return u1
			} else {
				t := u0
				u0 = u1
				u1 = u1 + t

				return u1
			}
		}
	}
}

// Fibotest : testing fibonaci with closures
func Fibotest() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
