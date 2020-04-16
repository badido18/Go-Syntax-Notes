package syntax

import "fmt"

func TakeInput() string {
	var name string
	fmt.Scanf("%s", &name)
	return name
}
