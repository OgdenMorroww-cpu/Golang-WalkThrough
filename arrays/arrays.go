package arrays

import "fmt"

func MyArrays() {
	myNumbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0

	for numbers := range myNumbers {

		sum += numbers
	}

	fmt.Printf("Sum: %v\n", sum)
	fmt.Printf("MyNumbers: %v\n", myNumbers)
}
