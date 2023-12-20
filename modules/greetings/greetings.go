package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name not provided")
	}
	return fmt.Sprintf("hello, %v. Welcome!", name), nil
}
