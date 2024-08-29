package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

var in = bufio.NewReader(os.Stdin)

type menu []menuItem
type dumbtype string



func (m menu) printitem() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}

}

func (m *menu) additem() {
	fmt.Println("Enter an new Menu Item")
	item, _ := in.ReadString('\n')
	*m = append(*m, menuItem{name: item, prices: make(map[string]float64)})

}

func PrintItem() {
	Data.printitem()
}

func AddItem() {
	Data.additem()
}
