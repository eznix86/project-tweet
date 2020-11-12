package main

import (
	"fmt"

	"github.com/theArtechnology/project-tweet/src/entity"
	"github.com/theArtechnology/project-tweet/src/repository"
)

func main() {
	u := repository.UserRepository{}
	user := entity.User{
		ID:        10,
		FirstName: "test",
	}
	user2 := entity.User{
		ID:        11,
		FirstName: "test1",
	}
	user3 := entity.User{
		ID:        12,
		FirstName: "test2",
	}
	u.Add(user)
	u.Add(user2)
	u.Add(user3)

	for _, user := range u.UserList {
		fmt.Println(user.FirstName)
	}
	usera, err := u.Find(0)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("found:")
	fmt.Println(usera.FirstName)

}
