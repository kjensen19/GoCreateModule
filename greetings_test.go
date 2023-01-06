package greetings

import (
	"regexp"
	"testing"
)

//TestHelloName calls greetings.Hello w/a name
//Checks for a valid return

func TestHelloName(t *testing.T) {
	name := "Terry"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Terry")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Terry") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls with empty string to check for error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
