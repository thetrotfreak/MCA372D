package main

import (
	"nescafe/menu"
)

func main() {
	m := menu.Menu{
		Items: []menu.Item{
			menu.Item{
				Name: "Tea",
				Prices: map[string]float64{
					"Hot":  50,
					"Cold": 75,
					"Iced": 100,
				},
			},
		},
	}

	m.Print()
	m.Add()
	m.Print()
	menu.Print(m)
	i := menu.Item{
		Name:   "Coffee",
		Prices: make(map[string]float64),
	}
	menu.Print(i)
}
