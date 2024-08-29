package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Printer interface {
	Print()
}

type Item struct {
	Name   string
	Prices map[string]float64
}

type Menu struct {
	Items []Item
}

func (m Menu) Print() {
	for _, item := range m.Items {
		item.Print()
	}
}

func (i Item) Print() {
	fmt.Println(i.Name)
	fmt.Println(strings.Repeat("-", len(i.Name)))
	for serving, price := range i.Prices {
		fmt.Printf("\t%16s%16.2f\n", serving, price)
	}
}

func Print(p Printer) {
	switch t := p.(type) {
	case Menu:
		t.Print()
	case Item:
		t.Print()
	default:
	}
}

func (m *Menu) Add() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Item Name? ")
	scanner.Scan()
	(*m).Items = append(
		m.Items,
		Item{
			Name:   scanner.Text(),
			Prices: make(map[string]float64),
		},
	)
}
