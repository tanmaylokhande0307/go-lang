package main

import (
	"fmt"
	"greetings"
)

func main(){
	message := greetings.Hello("tan")
	fmt.Println(message)
}