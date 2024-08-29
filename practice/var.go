package main

import "fmt"

type String string

func (str String) Join(iter ...string) string {
	if len(iter) == 0 {
		return str
	}

	for s := range iter {
	}
}

func main() {
	words := []string{
		"this",
		"is",
		"a",
		"sentence",
	}

	var sep String = " "

	fmt.Println(sep.Join(words...))
}
