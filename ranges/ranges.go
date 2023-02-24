package ranges

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func MyRanges() {

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	myNumbers := make([]int, 10)

	for i := range myNumbers {
		myNumbers[i] = 1 << uint(i)
	}

	for _, value := range myNumbers {
		fmt.Printf("%d\n", value)
	}

	countries := map[string]int{
		"Nigeria":       234587,
		"United States": 345868,
		"China":         120938457884,
		"India":         112378590057,
	}

	for key, value := range countries {
		fmt.Printf("%s has %d inhabitants\n", key, value)
	}
}
