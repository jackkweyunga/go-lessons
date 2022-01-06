package main

import (
	"fmt"

	"example.com/greetings"
)


func main() {
	// greet some great person of yours. Use the greetings module
	greeting := fmt.Println(greetings.Hello("Jackson"))
	return greeting
}



