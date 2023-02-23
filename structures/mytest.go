package structures

import "fmt"

type YourUser struct {
	Id             int
	Name, Location string
}

func (yourUser YourUser) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", yourUser.Name, yourUser.Location)
}

type YourPlayer struct {
	user   YourUser
	GameId int
}

func GreetingsForPlayers(yourPlayer YourPlayer) string {
	return yourPlayer.user.Greetings()
}
