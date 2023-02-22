package printing

import "fmt"

func Printing() {
	clyonModel := 12
	fmt.Println(clyonModel)

	name := "Caprica-Six"
	aka := fmt.Sprintf("Number %d", 6)
	fmt.Printf("%s is also known as %s\n", name, aka)

	planetName := "Jupiter"
	detailedPosition := fmt.Sprintf("Position Is Number %d", 5)
	fmt.Printf("%s  %s\n", planetName, detailedPosition)

}
