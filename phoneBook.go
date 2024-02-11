package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
)

type person struct {
	name    string
	surname string
	phone   string
}

var phoneBook []person

func search(phone string) *person {
	for _, pers := range phoneBook {
		if pers.phone == phone {
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

func getString(len int) string {
	temp := ""
	startChar := "!"

	for i := 0; i < len; i++ {
		myRand := rand.Intn(26) + 65
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
	}

	return temp
}

func populate(n int) {
	for i := 0; i < n; i++ {
		phoneBook = append(phoneBook, person{
			name:    getString(rand.Intn(10)),
			surname: getString(rand.Intn(10)),
			phone:   strconv.Itoa(rand.Intn(899999999) + 100000000),
		})
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		exe := path.Base(args[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
	}

	populate(10)

	/*phoneBook = append(phoneBook, person{
		name:    "Alex",
		surname: "Fern",
		phone:   "324809723",
	})
	phoneBook = append(phoneBook, person{
		name:    "Jean",
		surname: "Morningstar",
		phone:   "234809384",
	})*/

	switch args[1] {
	case "search":
		if len(args) != 3 {
			fmt.Printf("Usage: search Phone")
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
