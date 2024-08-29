package user

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestGetOne(t *testing.T) {
	expect := User{
		ID:       42,
		Username: "mrobot",
	}
	users = []User{expect}

	got, err := getOne(expect.ID)

	if err != nil {
		t.Fatal(err)
	}
	if got != expect {
		t.Errorf("did not get expected user. Got %+v, expected %+v", got, expect)
	}
}

func TestSlowOne(t *testing.T) {
	t.Parallel()
	t.Skip("skipped")
	time.Sleep(1 * time.Second)
}

func TestSlowTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func Example() {
	users = []User{
		{ID: 1, Username: "mrobot"},
	}
	u, err := getOne(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(u.ID, u.Username)

	// Output:
	// 1 mrobot
}
