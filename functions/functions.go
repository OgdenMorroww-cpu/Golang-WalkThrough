package functions

import "fmt"

func Functions() {
	result := getEven(11)
	fmt.Printf("Result: %v\n", result)

	answer := add(10, 45)
	fmt.Printf("answer: %v\n", answer)

	region, continent := location("Houston")
	fmt.Printf("Alisha lives in %s, %s\n", region, continent)

	getRole("Mark", "No code")
}

func getEven(myNumbers int) int {

	if myNumbers%2 == 0 {
		fmt.Printf("%v is an even numbers\n", myNumbers)
	} else {
		fmt.Printf("%v is an odd number\n", myNumbers)
	}

	return myNumbers
}

func add(x, y int) int {

	answer := x + y

	return answer
}

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "Houston", "Dallas", "Austin":
		region, continent = "Texas", "North America"
	case "Wuse II", "Central Park", "Maitama":
		region, continent = "Abuja", "Africa"
	default:
		region, continent = "Unknown", "Unknown"

	}
	return region, continent
}

func getRole(name string, role string) {
	switch role {
	case "Admin":
		fmt.Printf("%v get access to full website\n", name)
	case "Sub-Admin":
		fmt.Printf("%v get access to edit course\n", name)
	case "Testprep":
		fmt.Printf("%v get access to testprep\n", name)
	case "User":
		fmt.Printf("%v get to consume contents\n", name)
	default:
		fmt.Printf("%v is on trial\n", name)
	}
}
