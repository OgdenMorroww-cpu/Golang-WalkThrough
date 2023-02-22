package variables

import "fmt"

var (
	name     string
	age      int
	isOnline bool
)

func Variables() {
	name = "Ashley"
	fmt.Printf("Name: %v\n", name)
	age = 25
	fmt.Printf("Age: %v\n", age)
	isOnline = false
	fmt.Printf("Online Status: %v\n", isOnline)

	userName, planet, networth, age := "Elon", "Earth", 181, 54
	fmt.Printf("Details: %v %v %v %v\n", userName, planet, networth, age)
}
