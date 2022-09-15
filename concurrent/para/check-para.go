package main

import "time"

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTmCost = 180
)

func idCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	println("\tgoroutine-", id, ": idCheck ok\n")
	return idCheckTmCost
}

func bodyCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	println("\tgoroutine-", id, ": bodyCheck ok\n")
	return bodyCheckTmCost
}

func xRayCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	println("\tgoroutine-", id, ": xRayCheck ok\n")
	return xRayCheckTmCost
}

func airportCheck(id int) int {
	println("airportCheck ...")
	total := 0

	total += idCheck(id)
	total += bodyCheck(id)
	total += xRayCheck(id)

	println("\tgoroutine-", id, ": airportCheck ok\n")
	return total
}

func start(id int, f func(int) int, queue <-chan struct{}) <-chan int {
	c := make(chan int)
	go func() {
		total := 0
		for {
			_, ok := <-queue
			if !ok {
				c <- total
				return
			}
			total += f(id)
		}
	}()
	return c
}

func max(args ...int) int {
	n := 0
	for _, v := range args {
		if v > n {
			n = v
		}
	}
	return n
}

func main() {
	total := 0
	passengers := 30
	c := make(chan struct{})
	c1 := start(1, airportCheck, c)

	c2 := start(2, airportCheck, c)

	c3 := start(3, airportCheck, c)

	for i := 0; i < passengers; i++ {
		c <- struct{}{}
	}
	close(c)

	total = max(<-c1, <-c2, <-c3)
	println("total time cost:", total)
}
