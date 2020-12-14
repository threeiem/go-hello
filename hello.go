package main

import (
	"fmt"

	"rsc.io/quote"
	"github.com/threeiem/go-greeting"
)

var message string

func main() {
	fmt.Println(quote.Go())
	message = greeting.Hello("Unknown")
	fmt.Println(message)
}
