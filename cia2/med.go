package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Patient struct {
	name    string
	age     uint8
	gender  string
	contact string
	date    string
}

func menu() uint {
	fmt.Println()
	fmt.Println("---Medical Menu---")
	fmt.Println("1) Add a new patient")
	fmt.Println("2) Update patient details")
	fmt.Println("3) Schedule appointment")
	fmt.Println("4) Generate medical report")
	fmt.Println("5) Quit")

	var option uint
	fmt.Print("Choose > ")
	fmt.Scan(&option)

	return option
}

func newPatient(msg string) Patient {
	s := bufio.NewScanner(os.Stdin)
	var p Patient

	fmt.Println()
	fmt.Printf("---%v---\n", msg)
	fmt.Print("Name   ?    ")
	s.Scan()
	p.name = s.Text()
	fmt.Print("Age    ?    ")
	s.Scan()
	age, err := strconv.ParseUint(s.Text(), 10, 8)
	if err != nil {
		fmt.Println("You typed an invalid age, please redo.")
		return p
	} else {
		p.age = uint8(age)
	}
	fmt.Print("Gender ?    ")
	s.Scan()
	p.gender = s.Text()
	fmt.Print("Contact?    ")
	s.Scan()
	p.contact = s.Text()
	fmt.Print("Date   ?    ")
	s.Scan()
	p.date = s.Text()

	return p
}

func updatePatient(p *Patient) {
	n := newPatient("Patient Update")
	*p = n
}

func appointPatient(p *Patient) {
	var date string
	fmt.Print("Appointment date? ")
	fmt.Scanln(&date)
	p.date = date
}

func generateMedicalReport(p Patient) {
	fmt.Println("---Medical Report---")
	fmt.Println("Patient is", p.name)
	fmt.Println("Aged", p.age)
	fmt.Println("Is a", p.gender)
	fmt.Println("Contact via", p.contact)
	fmt.Println("Has an appointment on", p.date)
}

func main() {
	var p []Patient
loopmenu:
	o := menu()
	switch o {
	case 1:
		n := newPatient("Patient Add")
		p = append(p, n)
		goto loopmenu
	case 2:
		updatePatient(&p[len(p)-1])
		goto loopmenu
	case 3:
		appointPatient(&p[len(p)-1])
		goto loopmenu
	case 4:
		generateMedicalReport(p[len(p)-1])
		goto loopmenu
	case 5:
		fmt.Println("Exited!")
	default:
		fmt.Println("Only options are 1, 2, 3, 4, 5")
		goto loopmenu
	}
}
