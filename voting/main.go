package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Voter struct {
	Id   int
	Name string
	Age  byte
	Cast bool
}

type Candidate struct {
	Id          int
	Name        string
	Age         byte
	Affiliation Affiliations
	Votes       int
}

type Affiliations int

const (
	_ Affiliations = iota
	Foo
	Bar
	Baz
)

var (
	ErrNoName      = errors.New("Invalid Name")
	ErrNoAge       = errors.New("Invalid Age")
	ErrNoId        = errors.New("Invalid ID")
	ErrNoMenu      = errors.New("Invalid Menu Option")
	ErrNoAffil     = errors.New("Invalid Affiliation")
	ErrVoterExists = errors.New("Voter already registered")
	ErrCandExists  = errors.New("Candidate already registered")
)

func RUID() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

var (
	voters          = make(map[int]Voter)
	candidates      = make(map[int]Candidate)
	ballot          = make(map[int]int)
	nextVoterId     = RUID()
	nextCandidateId = RUID()
)

var menu = map[int]string{
	0: "Quit",
	1: "Register Voter",
	2: "Register Candidate",
	3: "Display Voters",
	4: "Display Candidates",
	5: "Delete Voter",
	6: "Delete Candidate",
	7: "Vote",
	8: "Voting Result",
}

type VotingSystem interface {
	Register() error
	Display()
	Delete() error
}

func Menu() (int, error) {
	fmt.Println("===Voting System===")
	for mo, mi := range menu {
		fmt.Println(mo, ")", mi)
	}

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return 0, ErrNoMenu
	}

	input := scanner.Text()
	option, err := strconv.Atoi(input)
	if err != nil {
		return option, ErrNoMenu
	}

	return option, nil
}

// method set for Voter
func (v *Voter) Register() error {
	fmt.Print("Enter voter name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return ErrNoName
	}
	v.Name = scanner.Text()

	fmt.Print("Enter voter age: ")
	if !scanner.Scan() {
		return ErrNoAge
	}
	age, err := strconv.Atoi(scanner.Text())
	if err != nil || age < 18 {
		return ErrNoAge
	}
	v.Age = byte(age)
	v.Id = nextVoterId()

	if _, exists := voters[v.Id]; exists {
		return ErrVoterExists
	}
	voters[v.Id] = *v

	return nil
}

func (v *Voter) Display() {
	fmt.Printf("Voter ID: %d, Name: %s, Age: %d, Cast Vote: %t\n", v.Id, v.Name, v.Age, v.Cast)
}

func (v *Voter) Delete() error {
	fmt.Print("Enter voter ID to delete: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return ErrNoId
	}
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return ErrNoId
	}

	if _, exists := voters[id]; !exists {
		return ErrNoId
	}

	delete(voters, id)
	return nil
}

// method set for Candidate
func (c *Candidate) Register() error {
	fmt.Print("Enter candidate name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return ErrNoName
	}
	c.Name = scanner.Text()

	fmt.Print("Enter candidate age: ")
	if !scanner.Scan() {
		return ErrNoAge
	}
	age, err := strconv.Atoi(scanner.Text())
	if err != nil || age < 18 {
		return ErrNoAge
	}
	c.Age = byte(age)

	fmt.Print("Enter candidate affiliation (0: Foo, 1: Bar, 2: Baz): ")
	if !scanner.Scan() {
		return ErrNoAffil
	}
	affil, err := strconv.Atoi(scanner.Text())
	if err != nil || affil < 0 || affil > 2 {
		return ErrNoAffil
	}
	c.Affiliation = Affiliations(affil)
	c.Id = nextCandidateId()

	if _, exists := candidates[c.Id]; exists {
		return ErrCandExists
	}
	candidates[c.Id] = *c

	return nil
}

func (c *Candidate) Display() {
	fmt.Printf("Candidate ID: %d, Name: %s, Age: %d, Affiliation: %v, Votes: %d\n", c.Id, c.Name, c.Age, c.Affiliation, c.Votes)
}

func (c *Candidate) Delete() error {
	fmt.Print("Enter candidate ID to delete: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return ErrNoId
	}
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return ErrNoId
	}

	if _, exists := candidates[id]; !exists {
		return ErrNoId
	}

	delete(candidates, id)
	return nil
}

func Vote() error {
	fmt.Print("Enter voter ID: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return ErrNoId
	}
	voterId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return ErrNoId
	}

	voter, exists := voters[voterId]
	if !exists {
		return ErrNoId
	}

	if voter.Cast {
		return errors.New("Voter has already cast their vote")
	}

	fmt.Print("Enter candidate ID to vote for: ")
	if !scanner.Scan() {
		return ErrNoId
	}
	candidateId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return ErrNoId
	}

	candidate, exists := candidates[candidateId]
	if !exists {
		return ErrNoId
	}

	candidate.Votes++
	candidates[candidateId] = candidate
	voter.Cast = true
	voters[voterId] = voter

	return nil
}

func Result() {
	fmt.Println("===Voting Results===")
	maxVotes := 0
	var winnerCandidate *Candidate
	for _, c := range candidates {
		if c.Votes > maxVotes {
			winnerCandidate = &c
		}
	}
	if winnerCandidate != nil {
		winnerCandidate.Display()
		fmt.Println(winnerCandidate.Name, "won by", winnerCandidate.Votes, "vote(s)")
	}
}

func main() {
	for {
		option, err := Menu()
		if err != nil {
			fmt.Println("Please Retry")
			continue
		}

		switch option {
		case 0:
			fmt.Println("Quitting...")
			return
		case 1:
			voter := Voter{}
			if err := voter.Register(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Voter registered successfully")
			}
		case 2:
			candidate := Candidate{}
			if err := candidate.Register(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Candidate registered successfully")
			}
		case 3:
			for _, voter := range voters {
				voter.Display()
			}
		case 4:
			for _, candidate := range candidates {
				candidate.Display()
			}
		case 5:
			voter := Voter{}
			if err := voter.Delete(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Voter deleted successfully")
			}

		case 6:
			candidate := Candidate{}
			if err := candidate.Delete(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Candidate deleted successfully")
			}
		case 7:
			Vote()
		case 8:
			Result()
		default:
		}
	}
}
