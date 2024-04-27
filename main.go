package main

import (
	"log"

	"github.com/ErtelDG/user-service/api"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	userById, err := api.GetUserbyId(1)
	check(err)
	log.Println(userById)
}
