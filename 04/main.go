// simplifying repetitive error handling
// resources: https://go.dev/blog/error-handling-and-go
package main

import (
	"errors"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/givemeerror", errorGiver)

	http.ListenAndServe(":8090", nil)
}

func errorGiver(w http.ResponseWriter, req *http.Request) {
	_, err := os.Open("filename.txt")

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if err := maybeErrorHandler(true); err != nil {
		log.Fatal(err.Error())
		return
	}

	return
}

func maybeErrorHandler(decided bool) error {
	if decided == true {
		return errors.New("You decided to fuckup.")
	}

	return nil
}
