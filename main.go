package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	r, err := doError()
	if err != nil {
		log.Printf("There was an error: %v\n", err)
	}
	fmt.Println("My message:", r)

	r, _ = doNoError()
	fmt.Println(r)

	err = moreError()
	if err != nil {
		log.Println(err)
	}

}

func doError() (string, error) {
	return "", errors.New("This is my custom error")
}

func doNoError() (string, error) {
	return "My response", nil
}

func moreError() error {
	errCode := 401
	return fmt.Errorf("This is my custom error from frmt package: %v", errCode)
}
