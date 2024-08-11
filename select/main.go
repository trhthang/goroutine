package main

import "time"

// googleSearch
func googleSearch(result chan string) {

	time.Sleep(3 * time.Second)

	result <- "found from Google"
}

// bingSearch
func bingSearch(result chan string) {

	time.Sleep(4 * time.Second)

	result <- "found from Bing"
}

func main() {

	chanGoogle := make(chan string)
	chanBing := make(chan string)

	go googleSearch(chanGoogle)
	go bingSearch(chanBing)

}
