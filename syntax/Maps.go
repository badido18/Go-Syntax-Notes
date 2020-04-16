package syntax

import "fmt"

func TestMaps() {

	var Names = [5]string{"Moh", "Bernard", "Asperge", "Bob", "Banana"}
	Ages := []int{20, 21, 23, 24, 25}
	//maps
	Persons := make(map[string]int)
	Persons["madjid"] = 20
	for i := 0; i < len(Names); i++ {
		Persons[Names[i]] = Ages[i]
	}
	fmt.Println(Persons)
	// a tab in a map keys are ids and (uint == int >0)
	tmp := make(map[uint][]int)
	tmp[1] = Ages
	fmt.Println(tmp)

	delete(Persons, "Asperge")
	fmt.Println(Persons)

	infos := map[string]int{"A": 0x41}
	fmt.Println(infos)
}
