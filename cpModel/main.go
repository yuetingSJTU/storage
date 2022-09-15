package main

import "fmt"

func Producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func Consumer(id int, ch chan int, done chan bool) {
	for {
		value, ok := <-ch
		if ok {
			fmt.Println("id:", id, ", recv:", value)
		} else {
			fmt.Println("id:", id, ", closed")
			break
		}
	}
	done <- true
}

func main() {
	ch := make(chan int, 3)

	coNum := 2
	done := make(chan bool, coNum)
	for i := 0; i < coNum; i++ {
		go Consumer(i, ch, done)
	}

	go Producer(ch)

	for i := 0; i < coNum; i++ {
		<-done
	}
}
