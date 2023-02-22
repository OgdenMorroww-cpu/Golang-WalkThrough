package packages

import (
	"fmt"
	"math"
	"math/rand"
)

func Packages() {
	fmt.Println("My favourite number is", rand.Intn(50))

	randomNumbers := rand.Intn(100)

	for i := 0; i <= randomNumbers; i++ {
		fmt.Printf("Numbers: %v\n", i)
	}

	counters := math.Sqrt(100)

	fmt.Printf("Counters : %v\n", counters)
}
