package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	id    uint
	name  string
	age   byte
	grade string
}

// a poorly written and perhaps a poor choice
// closure to mimic AUTO_INCREMENT of a databse
func nextID() func() uint {
	var i uint = 0
	return func() uint {
		i++
		return i
	}
}

func NewStudent(idfunc func() uint) *Student {
	var s Student
	scanner := bufio.NewScanner(os.Stdin)

	s.id = idfunc()

	fmt.Print("Name  : ")
	scanner.Scan()
	s.name = scanner.Text()

	fmt.Print("Age   : ")
	// Go took inspirations from C
	// and with it, took its crap input mechanism, imo
	// mixing fmt.Scant & bufio.Scanner
	// is a big No-No!
	scanner.Scan()
	age, _ := strconv.ParseUint(scanner.Text(), 10, 8)
	s.age = byte(age)

	fmt.Print("Grade : ")
	scanner.Scan()
	s.grade = scanner.Text()

	return &s
}

// func main() {
// 	var totalstudent int

// 	// stduent information system
// 	// I want a stuident_id:student k:v
// 	infosys := make(map[uint]Student)

// 	// we need this same closure
// 	// pass it around to maintin that AUOT_INCREMENT'ness
// 	autoid := nextID()

// 	fmt.Print("Number of students? ")
// 	fmt.Scanf("%d", &totalstudent)

// 	for i := 1; i <= totalstudent; i++ {
// 		s := NewStudent(autoid)
// 		infosys[s.id] = *s
// 	}

// 	fmt.Println(infosys)
// }
