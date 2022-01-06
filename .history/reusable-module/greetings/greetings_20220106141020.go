

package greetings


import "fmt"

// returns a greeting for the named person
func Hello(name string) string {
	// returns a greeting with the name inn it.
	message := fmt.Sprintf("Hi %v. Welcome! ", name)
	return message
}
