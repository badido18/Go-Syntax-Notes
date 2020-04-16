package main

import (
	"fmt"

	"./syntax"
)

//normal function

func DisplayHello(n string) string {
	return "Hello " + n
}

//main function

func main() {

	name := syntax.TakeInput()
	//Arrays
	var Names = [5]string{"Moh", "Bernard", "Asperge", "Bob", "Banana"}
	Ages := []int{20, 21, 23, 24, 25}
	fmt.Println(Names[1:3])
	fmt.Println(Ages[2:5])

	//Conditionals
	if "madjid" == name {
		fmt.Println("he is in !")
	} else if name != "moh" {
		fmt.Println("he is not in :( ")
	}
	switch name {
	case "madjid":
		fmt.Println("Yeah", name)
	case "moh":
		fmt.Println("Yeah", name)
	default:
		fmt.Println("not known")
	}

	//Loops
	for i := 0; i < len(Names); i++ {
		fmt.Printf("%s \n", DisplayHello(Names[i]))
	}

	// pointers like C very useful
	num := 10
	address := &num //a pointer
	fmt.Println(num, address)
	fmt.Println(*address, *&num)

}
