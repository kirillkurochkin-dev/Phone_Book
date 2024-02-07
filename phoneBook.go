package main

import (
	"fmt"
	"os"
	"path"
)

type person struct {
	name    string
	surname string
	phone   string
}

var phoneBook []person

func search(surname string) *person {
	for _, pers := range phoneBook {
		if pers.surname == surname {
			return &pers
		}
	}
	return nil
}

func list() {
	for _, p := range phoneBook {
		fmt.Println(p)
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		exe := path.Base(args[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
	}

	phoneBook = append(phoneBook, person{
		name:    "Alex",
		surname: "Fern",
		phone:   "324809723",
	})
	phoneBook = append(phoneBook, person{
		name:    "Jean",
		surname: "Morningstar",
		phone:   "234809384",
	})

	switch args[1] {
	case "search":
		if len(args) != 3 {
			fmt.Printf("Usage: search Username")
			return
		}
		pers := search(args[2])
		if pers == nil {
			fmt.Println("Person is not found: ", args[2])
			return
		}
		fmt.Println(*pers)
	case "list":
		list()
	default:
		fmt.Println("Option is not valid")
	}

}
