package arrays

import "fmt"

func MyArrays() {
	var a [2]string
	a[0] = "Mercury"
	a[1] = "Venus"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	messages := [2]string{"Hello", "World"}

	fmt.Printf("Messages: %v\n", messages)

	example := [...]string{"Pass", "Value", "Down"}

	fmt.Printf("Examples %v\n", example)

	quotes := [...]string{"A wise", "Sage", "Once", "Said"}

	fmt.Printf("Quotes %q\n", quotes)

	var myA [2][3][4]string

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				myA[i][j][k] = fmt.Sprintf("row %d - column %d - %d", i+1, j+1, k+1)
			}
		}
	}
	fmt.Printf("%q\n", myA)

}
