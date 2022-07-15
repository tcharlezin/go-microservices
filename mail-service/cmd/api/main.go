package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const webPort string = "80"

func main() {
	app := Config{}

	log.Println("Starting mail service on port ", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprint(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
