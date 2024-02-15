package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Person struct {
	Name       string
	Surname    string
	Phone      string
	LastAccess string
}

type PhoneBook []Person

func (p PhoneBook) Len() int {
	return len(p)
}

func (p PhoneBook) Less(i, j int) bool {
	if p[i].Surname == p[j].Surname {
		return p[i].Name < p[j].Name
	}
	return p[i].Surname < p[j].Surname
}

func (p PhoneBook) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var phoneBook PhoneBook
var index map[string]int

const CSVFILE string = ""

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := Person{
			Name:       line[0],
			Surname:    line[1],
			Phone:      line[2],
			LastAccess: line[3],
		}
		phoneBook = append(phoneBook, temp)
	}
	return nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range phoneBook {
		temp := []string{row.Name, row.Surname, row.Phone, row.LastAccess}
		err := csvwriter.Write(temp)
		if err != nil {
			return err
		}
	}
	csvwriter.Flush()
	return nil
}

func createIndex() error {
	index = make(map[string]int)
	for i, k := range phoneBook {
		key := k.Phone
		index[key] = i
	}
	return nil
}

func insert(person *Person) error {
	_, ok := index[(*person).Phone]
	if ok {
		return fmt.Errorf("%s already exists", person)
	}
	phoneBook = append(phoneBook, *person)

	_ = createIndex()

	err := saveCSVFile(CSVFILE)

	if err != nil {
		return err
	}
	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	phoneBook = append(phoneBook[:i], phoneBook[i+1:]...)
	delete(index, key)

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func search(phone string) *Person {
	for _, pers := range phoneBook {
		if pers.Phone == phone {
			return &pers
		}
	}
	return nil
}

func list() {
	sort.Sort(phoneBook)
	for _, p := range phoneBook {
		fmt.Println(p)
	}
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

func setCSVFILE() error {
	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("Creating", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	fileInfo, err := os.Stat(CSVFILE)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILE, "not a regular file!")
		return
	}

	err = readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}

	switch args[1] {
	case "insert":
		if len(args) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(args[4], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
	case "delete":
		if len(args) != 3 {
			fmt.Println("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(args[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		err := deleteEntry(t)
		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if len(args) != 3 {
			fmt.Println("Usage: search Number")
			return
		}
		t := strings.ReplaceAll(args[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := search(t)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}
		fmt.Println(*temp)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}

}
