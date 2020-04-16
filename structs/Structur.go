package structs

import "fmt"

//a struct

type Bike struct {
	brand string
	typ   string
	age   int
}

//methods

func (bike Bike) DisplayAge() {
	fmt.Printf("its age : %d months\n", bike.age)
}

func (bike *Bike) DecremAge() {
	bike.age--
}

//interface

type Caracteristics interface {
	speed() int
}

func (bike Bike) speed() int {
	fmt.Println("the speed is not defined hahaah", bike.brand)
	return 0
}

func TestStruct() {
	//Structs testing
	//look at the top to find it
	Mybike := Bike{brand: "GOsport", typ: "VTT", age: 1}

	fmt.Println(Mybike)

	Mybike = Bike{"ClimMountain", "offroad", 2}
	Mybike.age++

	//using methods defined on top

	Mybike.DisplayAge()
	Mybike.DecremAge()
	Mybike.DisplayAge()

	//using interfaces
	//this is a test an has no meaning for using interfaces
	Mybike.speed()

}
