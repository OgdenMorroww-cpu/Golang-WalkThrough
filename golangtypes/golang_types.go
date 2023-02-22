package golangtypes

import "fmt"

func Types() {

	canVote := 0

	if canVote < 18 {
		fmt.Println("Hey you can't vote now")
	} else if canVote >= 18 {
		fmt.Println("Yes you can vote")
	} else {
		fmt.Println("failed age is not entered")
	}
}
