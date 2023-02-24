package arrays

import "fmt"

func SliceExample() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(p)

	mySlice := []int{4, 6, 8, 10, 12, 14}
	fmt.Println(mySlice)

	fmt.Printf("Slice %v\n", mySlice[1:4])
	fmt.Printf("Slice %v\n", mySlice[:3])
	fmt.Printf("Slice %v\n", mySlice[4:])

	names := []string{
		"John",
		"Alison",
		"Diana",
		"Josh",
		"Dawn",
		"Rosia",
	}
	names = append(names, "Mylie", "drake", "Andre")
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	cities := make([]string, 5)
	cities[0] = "Talinn"
	cities[1] = "Bern"
	cities[2] = "Perth"
	cities[3] = "San francisco"
	cities[4] = "Vienna"

	fmt.Printf("%q", cities)

	todoTasks := make([]string, 5)

	todoTasks[0] = "Go to gymn"
	todoTasks[1] = "Cook"
	todoTasks[2] = "Buy bread"
	todoTasks[3] = "Walk my dogs"
	todoTasks[4] = "Go for run"

	for task, index := range todoTasks {
		fmt.Printf("Your task no: %v is %v\n", task, index)
	}

	actors := make([]string, 3)
	actors[0] = "Tom cruise"
	actors = append(actors, "Keanu Reeves", "Ryan Reynolds", "Ben Afleck", "Matt Damon")
	fmt.Printf("%q\n", actors)

	heroes := []string{"Batman", "Superman", "Wonder Woman"}
	alterEgoes := []string{
		"Bruce wayne",
		"Clark kent",
		"Diana prince",
	}
	heroes = append(heroes, alterEgoes...)
	fmt.Printf("%q\n", heroes)
	fmt.Println(len(heroes))

	cars := []string{}
	fmt.Println(len(cars))
	fmt.Println(cap(cars))

}
