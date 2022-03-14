package main

import (
	"errors"
	"log"
)

func main() {
	err := predefinedErrors()

	if err != nil {
		log.Fatalf(err.Error())
	}
}

func predefinedErrors() error {
	result := 1

	if result == 2 {
		return errors.New("Wrong result")
	}

	return nil
}
