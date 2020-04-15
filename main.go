package main

import (
	"fmt"
	"math"
	"./mypkgfuncs"
	"net/http"
)

//normal function

func DisplayHello(n string) string {
	return "Hello " + n
}

//closures

func Closure() func(string) int{
	Nmes := ""
	return func(a string) int{
		Nmes += a
		return len(Nmes)
	} 
}

//a struct 

type Bike struct {
	brand string
	typ string
	age int
}

//methods

func (bike Bike) DisplayAge() {
		fmt.Printf("its age : %d months\n",bike.age)
}

func (bike *Bike) DecremAge(){
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

//mywebserver function

func index (writer http.ResponseWriter , Request *http.Request){
	fmt.Fprintf(writer, "<H1>My Go Server <H1>")
}


//main function

func main() {

	//data types
	a := "this is a string"
	var b int = 5
	c := true
	d := 5.48
	const e float64 = 6e48

	//IO functions
	fmt.Println("Let's GO !", a, b, c, d+e)
	fmt.Printf("%T\n", e)
	fmt.Println(math.Max(math.Round(0.7), math.Floor(2.7)))
	fmt.Println(mypkgfuncs.Getpi())
	fmt.Print("Enter your name : ")

	//taking input
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println(DisplayHello(name))

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

	//maps

	Persons := make(map[string]int)
	Persons[name] = 20
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

	//ranges

	//for looping through maps

	for name, age := range Persons {
		fmt.Printf("%s is %d \n", name, age)
	}
	//either this way in an array
	for _, x := range Ages {
		fmt.Println(x)
	}

	// pointers like C very useful
	num := 10
	address := &num //a pointer
	fmt.Println(num,address)
	fmt.Println(*address , *&num)

	//Closure 
	//Look at the top to find the closure 
	na := Closure()
	fmt.Println(na("Madjid"))

	//Structs testing
	//look at the top to find it
	Mybike := Bike{brand : "GOsport" , typ : "VTT" , age : 1}

	fmt.Println(Mybike)

	Mybike = Bike{ "ClimMountain","offroad",2}
	Mybike.age++

	//using methods defined on top

	Mybike.DisplayAge()
	Mybike.DecremAge()
	Mybike.DisplayAge()

	//using interfaces
	//this is a test an has no meaning for using interfaces
	Mybike.speed()

	//web test
	//running a web server


	PORT := "3000"
	http.HandleFunc("/",index)
	fmt.Println("Starting at Port :",PORT)
	http.ListenAndServe(":"+PORT,nil)



}
