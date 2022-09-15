package main

import "time"

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTmCost = 180
)

func idCheck() int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	println("\tidcheck ok")
	return idCheckTmCost
}

func bodyCheck() int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	println("\tbodycheck ok")
	return bodyCheckTmCost
}

func xRayCheck() int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	println("\txRaycheck ok")
	return xRayCheckTmCost
}

func airportCheck() int {
	println("airportCheck ...")
	total := 0

	total += idCheck()
	total += bodyCheck()
	total += xRayCheck()

	println("airportCheck ok")
	return total
}

func main() {
	total := 0
	passengers := 30
	for i := 0; i < passengers; i++ {
		total += airportCheck()
	}
	println("total time cost:", total)
}
