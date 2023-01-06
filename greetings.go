package greetings

import "fmt"

// Hello returns a greeting for the named person
// ⬇			⬇️param	type	⬇️return type
func Hello(name string) string {
	//⬆️Capital letter means a function can be called by a function not in the package(exported name)
	//return a greeting that embeds the name
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	//	⬆️shortcut to declare and intialize
	return message
}
