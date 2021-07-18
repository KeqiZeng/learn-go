package main

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	nanosecond := int64(time.Now().Nanosecond())
	rand.Seed(nanosecond)
	var wg sync.WaitGroup
	wg.Add(1)

	toSignalStation := make(chan message)
	toEarth := make(chan message)
	earthToStation := make(chan bool)
	mg := newMarsGrid()
	mg.lifeNumber()
	var findAll [roverNum + 1]chan bool
	for i := 0; i < roverNum+1; i++ {
		findAll[i] = make(chan bool)
	}
	var rover [roverNum]rover
	for i := 0; i < roverNum; i++ {
		rover[i] = newRover("Rover" + strconv.Itoa(i))
		go rover[i].randSearchLife(&mg, toSignalStation, findAll[i], findAll[i+1])
	}
	go func() {
		for {
			var searchNum int
			for i := 0; i < roverNum; i++ {
				searchNum += rover[i].searchNum
			}
			if searchNum == (width * height) {
				findAll[0] <- true
				return
			}
		}
	}()
	go func() {
		signalStation(toSignalStation, toEarth, earthToStation, findAll[roverNum])
		wg.Done()
	}()
	go earthStation(toEarth, earthToStation)
	wg.Wait()
}
