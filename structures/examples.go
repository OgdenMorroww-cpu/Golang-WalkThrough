package structures

import "fmt"

type Person struct {
	Name string
}

type Employee struct {
	Salary int
	Role   string
	Person
}

func doWork(employee Employee) {
	fmt.Printf("Hey My Name Is: %v\n", employee.Name)
	fmt.Println("And I'm Doing Work")
}

type MyUser struct {
	Id                 int
	MyName, MyLocation string
}

func (myUser *MyUser) Greetings() string {
	return fmt.Sprintf("Hi %s from %s\n", myUser.MyName, myUser.MyLocation)
}

type MyPlayer struct {
	MyUser
	GameId int
}

func Examples() {
	employee := Employee{
		Person: Person{Name: "Arash"},
		Salary: 4500,
		Role:   "Analyst",
	}
	fmt.Printf("Full Name: %v\n", employee.Name)
	fmt.Printf("Salary: %v\n", employee.Salary)
	fmt.Printf("Role: %v\n", employee.Role)
	doWork(employee)
	fmt.Println("")
	player := MyPlayer{
		MyUser: MyUser{Id: 42, MyName: "Justin", MyLocation: "LA"},
		GameId: 450,
	}
	fmt.Printf("UserID: %v\n", player.Id)
	fmt.Printf("UserName: %v\n", player.MyName)
	fmt.Printf("Location: %v\n", player.MyLocation)
	fmt.Printf("Game-ID: %v\n", player.GameId)
	fmt.Println(player.Greetings())
}
