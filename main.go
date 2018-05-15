package main

import (
	"github.com/wongoo/go-app-example/mongo"
	"github.com/wongoo/go-app-example/domain"
	"fmt"
)

func main() {
	user := domain.User{
		Id:   "1",
		Name: "wongoo",
		Sex:  1,
	}
	mongo.InsertCollection("user", user)
	fmt.Println("done")
}
