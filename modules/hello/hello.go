package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hellos([]string{"tan","lok"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
