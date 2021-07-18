package main

import (
	"math/rand"
	"time"
)

const roverNum = 5

type rover struct {
	name      string
	pos       point
	direction point
	searchNum int
}

func (r *rover) turnLeft() {
	r.direction = point{
		x: r.direction.y,
		y: -r.direction.x,
	}
}

func (r *rover) turnRight() {
	r.direction = point{
		x: -r.direction.y,
		y: r.direction.x,
	}
}

func (r *rover) stepOverOn() {
	if r.direction.x == 0 {
		r.direction.x = r.direction.y
	}
	if r.direction.y == 0 {
		r.direction.y = r.direction.x
	}
}

func (r *rover) stepOverOff() {
	if r.direction.x != 0 && r.direction.y != 0 {
		if i := rand.Intn(2); i == 0 {
			r.direction.x = 0
		} else {
			r.direction.y = 0
		}
	}
}

func (r *rover) move(mg *marsGrid) {
	if (r.pos.x+r.direction.x) >= width || (r.pos.x+r.direction.x) < 0 || (r.pos.y+r.direction.y) >= height || (r.pos.y+r.direction.y) < 0 || mg[r.pos.y+r.direction.y][r.pos.x+r.direction.x].occupied == true {
		switch n := rand.Intn(2); n {
		case 0:
			r.turnLeft()
		case 1:
			r.turnRight()
		}
		return
	}
	mg[r.pos.y][r.pos.x].mu.Lock()
	mg[r.pos.y][r.pos.x].occupied = false
	mg[r.pos.y][r.pos.x].mu.Unlock()
	mg[r.pos.y+r.direction.y][r.pos.x+r.direction.x].mu.Lock()
	r.pos.x += r.direction.x
	r.pos.y += r.direction.y
	mg[r.pos.y][r.pos.x].occupied = true
	if mg[r.pos.y][r.pos.x].searched == false {
		mg[r.pos.y][r.pos.x].searched = true
		r.searchNum++
	}
	mg[r.pos.y][r.pos.x].mu.Unlock()
}

func (r *rover) randSearchLife(mg *marsGrid, toSignalStation chan message, findAllR, findAllS chan bool) {
	var msg message
	var steps int
	switch n := rand.Intn(4); n {
	case 0:
		r.direction.x = -1
		r.direction.y = -0
	case 1:
		r.direction.x = 1
		r.direction.y = 0
	case 2:
		r.direction.x = 0
		r.direction.y = -1
	case 3:
		r.direction.x = 0
		r.direction.y = 1
	}
	keepFinding := time.After(time.Millisecond)
	for {
		select {
		case <-findAllR:
			findAllS <- true
			return
		case <-keepFinding:
			r.move(mg)
			steps++
			if mg[r.pos.y][r.pos.x].lifeValue >= 900 && mg[r.pos.y][r.pos.x].reported == false {
				msg = message{
					roverName: r.name,
					pos:       r.pos,
					lifeValue: mg[r.pos.y][r.pos.x].lifeValue,
				}
				toSignalStation <- msg
				mg[r.pos.y][r.pos.x].reported = true
			}
			if st := rand.Intn(10) + 1; steps%st == 0 {
				switch n := rand.Intn(4); n {
				case 0:
					r.turnLeft()
				case 1:
					r.turnRight()
				case 2:
					r.stepOverOn()
				case 3:
					r.stepOverOff()
				}
			}
			keepFinding = time.After(time.Millisecond)
		}
	}
}

func newRover(name string) rover {
	var r rover
	r.name = name
	r.pos.x = rand.Intn(width)
	r.pos.y = rand.Intn(height)
	return r
}
