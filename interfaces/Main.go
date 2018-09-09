package main

import (
	"fmt"
)

// Screamer is an interface for screaming at console
type Screamer interface {
	Scream(msg string, times int)
}

// ConsoleScreamer is an instance of Sreamer interface
type ConsoleScreamer struct{}

// Scream is the scream method for this
func (cs ConsoleScreamer) Scream(msg string, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("Screaming!!! %v\n", msg)
	}
}

// LoudScreamer is an instance of Sreamer interface
type LoudScreamer struct{}

// Scream is the scream method for this
func (lcs LoudScreamer) Scream(msg string, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("SCREAMING!!! %v\n", msg)
	}
}

func main() {
	fmt.Println("Hello, playground")

	var screamer Screamer = ConsoleScreamer{}

	screamer.Scream("Help!", 5)

	screamer = LoudScreamer{}

	screamer.Scream("HELP!", 5)
}
