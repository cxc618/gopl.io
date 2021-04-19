package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	message := fmt.Sprintf("Hi, %v\n", name)

	if name == "error" {
		return "", errors.New("empty name")
	}
	return message, nil
}
