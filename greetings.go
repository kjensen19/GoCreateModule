package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

	message := fmt.Sprintf(randomFormat(), name)
	//	⬆️shortcut to declare and intialize
	return message, nil
}

// init sets intial values for variables in the function
// init functions are executed at program start automatically, after global vars are intialized
func init() {
	rand.Seed(time.Now().UnixNano())
}

// ⬇️Lower case means it is only accessible in this package
func randomFormat() string {
	//A slice of messages
	//			⬇️ Non-specified size is what lets Go know this is a slice
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	//specifies a random index for the slice
	return formats[rand.Intn(len(formats))]
}
