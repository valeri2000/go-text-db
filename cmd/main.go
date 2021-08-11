package main

import (
	"fmt"

	db "github.com/valeri2000/go-text-db"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	database, err := db.NewDatabase("data.json")
	if err != nil {
		fmt.Println(err)
	}
	database.Put("0", "zero")
	database.Put("Ivan", Person{Name: "Ivan", Age: 15})

	var ivan Person
	temp, ok := database.Get("Ivan")
	if ok {
		ivan, ok = temp.(Person)
		if ok {
			fmt.Print(ivan)
		} else {
			fmt.Println("error retrieving Ivan")
		}
	} else {
		fmt.Println("error retrieving Ivan")
	}

	database.Print()
	defer database.Close()
}
