package structures

import "fmt"

type Location struct {
	Lat float64
	Lon float64
}

func Initializing() {
	x := new(Location)
	y := &Location{}
	fmt.Println(*x == *y)
}
