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

// Newew Hellos function calls the existing Hello function.
// This helps reduce duplication while also leaving both functions available.
// Hellos will return a map that associates each named person with a greeting
func Hellos(names []string) (map[string]string, error) {
	//Make the map syntax make(map[key-type]value-type)
	messages := make(map[string]string)
	//Loop through the slice of names, call function on each and check for missing values
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//in the map assoc retrieved message with name
		messages[name] = message
	}
	return messages, nil
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
