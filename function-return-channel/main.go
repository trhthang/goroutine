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

func main() {

	bridge := printSthing("hello")

	for i := 0; i <= 5; i++ {
		log.Println("receive from: ", <-bridge)
	}

}
