package syntax

import "fmt"

func TestRanges() {

	name := "madjid"
	var Names = [5]string{"Moh", "Bernard", "Asperge", "Bob", "Banana"}
	Ages := []int{20, 21, 23, 24, 25}

	Persons := make(map[string]int)
	Persons[name] = 20
	for i := 0; i < len(Names); i++ {
		Persons[Names[i]] = Ages[i]
	}
	//ranges

	//for looping through maps

	for name, age := range Persons {
		fmt.Printf("%s is %d \n", name, age)
	}
	//either this way in an array
	for _, x := range Ages {
		fmt.Println(x)
	}
}
