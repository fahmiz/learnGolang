package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	fmt.Println("running Error()")
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	fmt.Println("running run()")
	return &MyError{
		time.Now(),
		"Error Message: it didn't work",
	}
	// return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
