package main

import (
	"fmt"
	"log"

	"rsc.io/quote"
	"github.com/threeiem/go-greeting"
)

func main() {
	greeting, err := greet("Username")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting)
}

func greet(name string) (string, error) {
	log.SetPrefix("hello::greet(): ")
	log.SetFlags(0)

	return greeting.Hello(name)
}
