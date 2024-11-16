package main

import (
	"fmt"
	"roshpr.net/examples/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, from server!")
	fmt.Println(quote.Go())
	message := greetings.Hello("Glads")
	fmt.Println(message)
}
