package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nanosecond := time.Now().Nanosecond()
	rand.Seed(int64(nanosecond))
	var dog = Dog{name: "Lucky", food: "meat"}
	var cat = Cat{name: "Kitty", food: "fish"}
	var bear = Bear{name: "Wayney", food: "honey"}
	for day := 0; day < 3; day++ {
		fmt.Println("Day ", day)
		for clock := 0; clock < 24; clock++ {
			// in day
			m := rand.Intn(2)
			if clock > 8 && clock < 18 {
				switch n := rand.Intn(3); n {
				// dog
				case 0:
					if m == 0 {
						AnimalEat(dog)
					} else {
						AnimalMove(dog)
					}
				// cat
				case 1:
					if m == 0 {
						AnimalEat(cat)
					} else {
						AnimalMove(cat)
					}
				// bear
				case 2:
					if m == 0 {
						AnimalEat(bear)
					} else {
						AnimalMove(bear)
					}
				}
			} else {
				// in night
				switch n := rand.Intn(3); n {
				// dog
				case 0:
					AnimalSleep(dog)
				// cat
				case 1:
					AnimalSleep(cat)
				// bear
				case 2:
					AnimalSleep(bear)
				}
			}
		}

	}
}
