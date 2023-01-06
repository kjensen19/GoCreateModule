package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
//
//	⬇️param	type	⬇️return type
func Hello(name string) (string, error) {
	//⬆️Capital letter means a function can be called by a function not in the package(exported name)
	//return a greeting that embeds the name
	//If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	//	⬆️shortcut to declare and intialize
	return message, nil
}
