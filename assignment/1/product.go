package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Product struct {
	name  string
	price float64
}

type Store struct {
	store map[Product]uint
	// perhaps we can have a bool here
	// to let us know
	// whether the underlying
	// map is initialised or not
}

func (s Store) Update(p Product) {
	// whether prodcut is in store
	// does not matter
	// cause of zero value
	quantity := s.store[p]
	quantity++
	s.store[p] = quantity
}

func NewProduct() *Product {
	var p Product
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("PRODCUT")
	fmt.Print("\tName  : ")
	scanner.Scan()
	p.name = scanner.Text()
	fmt.Print("\tPrice : ")
	scanner.Scan()
	input, _ := strconv.ParseFloat(scanner.Text(), 64)
	p.price = input

	return &p
}

// func main() {
// 	var store Store
// 	// we have to make the map
// 	store.store = make(map[Product]uint)
// 	var prodcuts int

// 	fmt.Print("Product(s) to be added to the store? ")

// 	fmt.Scanf("%d", &prodcuts)
// 	for i := 1; i <= prodcuts; i++ {
// 		p := NewProduct()
// 		store.Update(*p)
// 	}

// 	fmt.Println(store.store)
// }
