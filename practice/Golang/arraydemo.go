package main

import (
	"bufio"
	"demo/menu"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println(menu.Data)
loop:
	for {
		fmt.Println("Please Enter your Choice")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add Item")
		fmt.Println("q)quit")

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.PrintItem()

		case "2":
			menu.AddItem()
		case "q":
			break loop
		default:
			fmt.Println("Please Enter the Correct Option")
		}
	}
}
