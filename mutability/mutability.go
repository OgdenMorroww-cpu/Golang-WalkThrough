package mutability

import "fmt"

type Artist struct {
	Name, Genre string
	Songs       int
}

func newReleases(a *Artist) int {

	a.Songs++

	return a.Songs
}

type Student struct {
	StudentName string
	Grade       int
}

func calculateScore(student *Student) int {
	student.Grade++
	return student.Grade
}

type MovieDirector struct {
	DirectorName    string
	TotalMoviesMade int
	Genre           string
}

func totalMoviesMade(movieDirector *MovieDirector) int {
	movieDirector.TotalMoviesMade++
	return movieDirector.TotalMoviesMade
}

func Mutability() {
	artist := &Artist{Name: "Taylor Swift", Genre: "Country", Songs: 44}
	fmt.Printf("%s released her %dth song\n", artist.Name, newReleases(artist))
	fmt.Printf("%s has a total of %d songs\n", artist.Name, artist.Songs)
	fmt.Println("")
	student := &Student{StudentName: "John", Grade: 50}
	fmt.Printf("%s original grade is %d points\n", student.StudentName, calculateScore(student))
	fmt.Printf("%s has a total grade of %d points\n", student.StudentName, student.Grade)
	fmt.Println("")
	directorName := &MovieDirector{
		DirectorName:    "Zack Snyder",
		TotalMoviesMade: 7,
		Genre:           "Action",
	}
	fmt.Printf("%s released his %dth movies\n", directorName.DirectorName, totalMoviesMade(directorName))
	fmt.Printf("%s has a total of %d movies\n", directorName.DirectorName, directorName.TotalMoviesMade)
}
