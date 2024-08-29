package main

import "fmt"

type Student struct {
	id    uint
	name  string
	age   byte
	grade rune
}

func nextID() func() int {
	i:= a
	func inner() int {
		i++
		return i
	}
}
		
func (s Student) SetID() {
}

func (s Student) NewStudent() Student {
	fmt.println("")
}

func main() {
	db := make(map[SID]Student)
	var s1 = Student{
		name:  "bivas kumar",
		age:   22,
		grade: 'A',
	}
	fmt.Println(s1)
}
