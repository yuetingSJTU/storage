// main.go
package main

import "log"

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	cal := new(Cal)
	var result Result
	cal.Square(11, &result)
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
