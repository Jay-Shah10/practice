package main

import (
	"github.com/Jay-Shah10/practice/go/hello-world/server"
	"log"
)

func main() {

	log.Println("Starting server.")
	server := server.New()
	server.Start()
}
