package main

import (
	"fmt"
	"math/rand"
	"time"
)

type message struct {
	roverName string
	pos       point
	lifeValue int
}

func signalStation(toSignalStation, toEarth chan message, earthToStation, findAllLastS chan bool) {
	msgs := make([]message, 0, 20)
	sendMsg := func() {
		ok := <-earthToStation
		if ok && len(msgs) != 0 {
			for _, m := range msgs {
				toEarth <- m
			}
			msgs = make([]message, 0, 20)
		}
	}
	keepSend := time.After(time.Millisecond)
	for {
		select {
		case <-findAllLastS:
			for {
				if len(msgs) == 0 {
					break
				}
				sendMsg()
			}
			close(toSignalStation)
			close(toEarth)
			fmt.Println("All lives have been found!")
			return
		case msg := <-toSignalStation:
			msgs = append(msgs, msg)
		case <-keepSend:
			sendMsg()
			keepSend = time.After(time.Millisecond)
		}
	}
}

func earthStation(toEarth chan message, earthToStation chan bool) {
	noMsg := time.After(300 * time.Millisecond)
	for {
		select {
		case msg := <-toEarth:
			fmt.Println(msg)
		case <-noMsg:
			if n := rand.Intn(10); n > 5 {
				earthToStation <- true
			} else {
				earthToStation <- false
			}
			noMsg = time.After(300 * time.Millisecond)
		}
	}
}
