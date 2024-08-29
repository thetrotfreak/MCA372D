package main

import (
	"bufio"
	"fmt"
	"os"
)

// struct field has an unrequired `availability`
type Book struct {
	title        string
	availability bool
}

// use a global for now
var Inventory map[string]bool = make(map[string]bool)

// a method of type Book should be implemented instead
func isBookAvailable(title string) bool {
	// looping is unnecessary
	// for t, a := range Inventory {
	// 	if strings.Compare(title, t) == 0
	// 		return a

	// existence check is mundane
	// cause of the Inventory
	// avail, err := Inventory[title]
	return Inventory[title]
}

func main() {
	Inventory["Learning Go"] = true
	Inventory["The C Programming Language"] = true
	Inventory["Java The Complete Reference"] = false

	var userquery string
	// fmt.Scanln(&userquery)
	// Scanln() wont work,
	// it needs to know what & how many to read
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Search for Book (title)? ")
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		fmt.Println("Something went wrong...")
	}

	userquery = scanner.Text()
	if isBookAvailable(userquery) == true {
		fmt.Printf("Book \"%v\" is available.", userquery)
	} else {
		fmt.Printf("Book \"%v\" is not available.", userquery)
	}
}
