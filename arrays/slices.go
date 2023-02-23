package arrays

import "fmt"

func MySlices() {
	results := oddNumbers([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21})
	fmt.Println(results)

	myNums := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Println(myNums)

	fmt.Println(myNums[1:4])
}

func oddNumbers(myOddNumber []int) int {
	sum := 0

	for index := range myOddNumber {
		sum += index
	}
	fmt.Printf("Sum: %v\n", sum)
	return sum
}
