package main

import (
	"fmt"
	"log"
)

func printSthing(msg string) chan string {
	result := make(chan string)

	go func() {
		for i := 0; i <= 5; i++ {
			result <- fmt.Sprintf("%s %d", msg, i)
		}
	}()

	return result
}

func fanIn(chan1, chan2 chan string) chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case <-chan1:
				c <- <-chan1
			case <-chan2:
				c <- <-chan2
			}

		}
	}()

	return c
}

func main() {

	coffee := printSthing("hello")

	bread := printSthing("bread order")

	serve := fanIn(coffee, bread)

	for i := 0; i <= 5; i++ {
		log.Println("receive from: ", <-serve)
	}

	log.Println("main finished")
}
