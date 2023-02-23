package structures

import "fmt"

type User struct {
	Id     int
	Name   string
	Region string
}

type Player struct {
	User   User
	GameId int
}

func Compositions() {
	player := Player{}
	player.User.Id = 42
	player.User.Name = "Matt"
	player.User.Region = "Mountian View"
	player.GameId = 90404
	fmt.Printf("%+v", player)
}
