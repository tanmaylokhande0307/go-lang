package greetings

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("hello, %v. Welcome!", name)
}
