package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type idea struct {
		Id         int
		Title      string
		Body       string
		User       int
		TimePosted time.Time
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	currentTime := time.Now()

	myIdea := idea{
		Id:         1,
		Title:      "Learn Go",
		Body:       "I should learn Go so that I know how to program with it.",
		User:       1,
		TimePosted: currentTime,
	}

	go func() {
		defer waitGroup.Done()

		fmt.Println("First function waiting for 10 seconds to come up with an idea...")
		time.Sleep(10 * time.Second)
		fmt.Println("Done with first function. Printing the idea.", myIdea)
	}()

	go func() {
		defer waitGroup.Done()

		fmt.Println("Done with the second function. ", time.Now())
	}()

	waitGroup.Wait()
}
