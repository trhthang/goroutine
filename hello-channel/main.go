package main

import (
	"log"
)

func printNumber(numChan chan int) {
	var result int
	for i := 0; i <= 100; i++ {
		result += i
	}
	numChan <- result
}

func printChar(strChan chan string) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		result += string(i)
	}
	strChan <- result
}

func main() {

	chanPrintNumber := make(chan int)

	chanPrintChar := make(chan string)

	go printNumber(chanPrintNumber)
	go printChar(chanPrintChar)

	log.Println("Kết quả từ printNumber: ", <-chanPrintNumber)
	log.Println("Kết quả từ printChar: ", <-chanPrintChar)

}

// Tại sao không cần WaitGroup?
// Channel Blocking: Khi bạn đọc từ một channel (sử dụng cú pháp <-channel), nếu không có giá trị nào có sẵn trong channel, lệnh này sẽ chặn (block) cho đến khi có giá trị được gửi vào channel. Điều này có nghĩa là chương trình sẽ dừng lại ở lệnh nhận giá trị từ channel cho đến khi giá trị đó có sẵn.
// Vì lý do này, khi bạn gọi <-chanPrintNumber, chương trình sẽ chặn tại đó cho đến khi goroutine printNumber hoàn thành và gửi kết quả vào chanPrintNumber. Tương tự, chương trình sẽ chặn tại <-chanPrintChar cho đến khi goroutine printChar gửi kết quả vào chanPrintChar.
// Không cần WaitGroup: Chính vì cơ chế chặn của channel, chương trình đảm bảo rằng các goroutine sẽ hoàn thành trước khi tiếp tục thực thi các lệnh tiếp theo, do đó không cần sử dụng WaitGroup để chờ các goroutine kết thúc.
