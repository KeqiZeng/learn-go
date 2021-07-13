package main

import (
	"fmt"
	"math/rand"
	"time"
)

var now = int64(time.Now().Nanosecond())

func main() {
	rand.Seed(now)
	const distance = 62100000 //km
	const secondsPerHour = 3600
	const hoursPerDay = 24
	var (
		speed      int // km/s
		days       int
		airways    string
		ticketType string
		price      int
	)
	fmt.Printf("%-15v %-18v %-15v %-10v\n", "太空航行公司", "飞行天数", "飞行类型", "价格(百万美元)")
	for i := 0; i < 10; i++ {
		// choose airways
		switch n := rand.Intn(3); n {
		case 0:
			airways = "Space Adventures"
		case 1:
			airways = "SpaceX"
		case 2:
			airways = "Virgin Galactic"
		}

		//speed
		speed = rand.Intn(15) + 16
		days = distance / (speed * secondsPerHour * hoursPerDay)

		// choose ticketType
		switch n := rand.Intn(2); n {
		case 0:
			ticketType = "单程"
		case 1:
			ticketType = "往返"
		}
		n := speed - 16
		price = 36 + n*1
		if ticketType == "往返" {
			price *= 2
		}
		fmt.Printf("%-25v %-20v %-20v %-20v\n", airways, days, ticketType, price)
	}
}
