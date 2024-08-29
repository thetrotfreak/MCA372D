package main

import (
	"io"
	"log"
	"net/http"
)

type Body struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

type Typicode interface {
	Post(url, body string) error
	Get(url string) error
}

func (b Body) Get(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	log.Println(string(body))
	return err
}

func main() {
	var b Body
	b.Get("https://jsonplaceholder.typicode.com/posts/1")
}
