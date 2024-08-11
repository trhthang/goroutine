package main

import (
	"fmt"
	"log"
	"sync"
)

func printNumber(wg *sync.WaitGroup) {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printChar(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func main() {
	// B1: Tạo 1 biến có kiểu sync.WaitGroup: wg
	var wg sync.WaitGroup

	// B2: wg.Add(số lượng goroutine cần phải đợi) : wg.add(2)

	wg.Add(2)

	// B3: ở mỗi goroutine, gọi method wg.Done() trước khi return

	// B4: gọi method wg.Wait() tại method mà cần các goroutine chạy xong (ví dụ hàm tại main)

	defer log.Print("finish")

	go printNumber(&wg)
	go printChar(&wg)

	wg.Wait()

}
