package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Printf("%-16v %5v %10v %5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Printf("=======================================\n")

	for count := 10; count > 0; count-- {
		// Spaceline options
		var spacelineIndex = rand.Intn(3)
		var spaceline = ""
		switch spacelineIndex {
		case 1:
			spaceline = "Space Adventures"
		case 2:
			spaceline = "SpaceX"
		default:
			spaceline = "Virgin Galactic"
		}

		// days
		var speed = rand.Intn(30-16+1) + 16 // km/s
		const distance = 62_100_000
		const secondsInMinute = 60
		const minutesInHour = 60
		const hoursInDay = 24
		var days = distance / (speed * secondsInMinute * minutesInHour * hoursInDay)

		// trip type
		var typeIndex = rand.Intn(2)
		var tripType = ""
		switch typeIndex {
		case 0:
			tripType = "One-way"
		default:
			tripType = "Round-trip"
		}

		// calculate cost
		var price = speed + 20
		if tripType == "Round-trip" {
			price *= 2
		}

		fmt.Printf("%-16v %5v %-10v $%5v\n", spaceline, days, tripType, price)
	}
}

