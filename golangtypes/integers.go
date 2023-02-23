package golangtypes

import (
	"fmt"
	"math/cmplx"
)

var (
	goIsFun         bool       = true
	maxInt          uint64     = 1<<64 - 1
	complexIntegers complex128 = cmplx.Sqrt(-5 + 12i)
)

func Integers() {
	const f = "%T(%v)\n"
	fmt.Printf(f, goIsFun, goIsFun)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complexIntegers, complexIntegers)

	region, country := locations("Frankurt")
	fmt.Printf("Tessy lives in %s %s\n", region, country)
}

func locations(cities string) (string, string) {
	var regions string
	var country string

	switch cities {
	case "LA", "San-francisco", "Palo alto":
		regions, country = "California", "United States"
	case "Berlin", "Frankurt", "Munich":
		regions, country = "Ulm", "Germany"
	case "Houston", "Austin", "Dallas":
		regions, country = "Texas", "United States"
	default:
		regions, country = "Unknown", "Unknown"
	}
	return regions, country
}
