package main

import (
	"net/http"
)

const PORT = ":8000"

type Config struct{}

func main() {
	app := Config{}

	srv := &http.Server{
		Addr:    PORT,
		Handler: app.Routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
