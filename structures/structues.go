package structures

import (
	"fmt"
	"time"
)

type BootCamp struct {
	Lat float64
	Lon float64

	Date time.Time

	IsNight bool
}

type Points struct {
	X, Y int
}

var (
	p = Points{1, 2}
	q = &Points{1, 2}
	r = Points{X: 1}
	s = Points{}
)

func Structures() {
	bootCamp := BootCamp{
		Lat:  34.012836,
		Lon:  -118.495338,
		Date: time.Now(),
	}

	fmt.Printf("Lat: %v\n", bootCamp.Lat)
	fmt.Printf("Lon: %v\n", bootCamp.Lon)
	fmt.Printf("Date: %v\n", bootCamp.Date)
	fmt.Printf("Day Status: %v\n", bootCamp.IsNight)

	fmt.Println(p, q, r, s)
}
