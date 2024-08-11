package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func countFirstFile(result chan int, filePath string, keywork string) {
	//logic tính số lần xuất hiện của từ khóa trong file
	var numOfOcc int
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println(err)
		result <- 0
		return
	}

	numOfOcc = strings.Count(string(fileContent), keywork)

	result <- numOfOcc
	defer close(result)
}

func countSecondFile(result chan int, filePath string, keywork string) {
	var numOfOcc int
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println(err)
		result <- 0
		return
	}

	numOfOcc = strings.Count(string(fileContent), keywork)

	result <- numOfOcc
	defer close(result)
}

func main() {

	countFirstChan := make(chan int)

	countSecondChan := make(chan int)

	go countFirstFile(countFirstChan, "1.txt", "Man City")
	go countSecondFile(countSecondChan, "2.txt", "Man City")

	log.Println("Tổng số lần xuất hiện trong 2 file: ", <-countFirstChan+<-countSecondChan)

}
