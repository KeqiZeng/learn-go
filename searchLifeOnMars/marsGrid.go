package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	width  = 100
	height = 100
)

type point struct {
	x int
	y int
}

type marsPoint struct {
	p         point
	lifeValue int
	occupied  bool
	reported  bool
	searched  bool
	mu        sync.Mutex
}
type marsGrid [height][width]marsPoint

//func (mg marsGrid) searchFinish() bool {
//for h := 0; h < height; h++ {
//for w := 0; w < width; w++ {
//if mg[h][w].lifeValue >= 900 {
//if mg[h][w].searched == false {
//return false
//}
//}
//}
//}
//return true
//}

func (mg marsGrid) lifeNumber() {
	var num int
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if mg[h][w].lifeValue >= 900 {
				num++
			}
		}
	}
	fmt.Println("lifeNumber is", num)
}

func newMarsGrid() marsGrid {
	var mg marsGrid
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			mg[h][w].p.x = w
			mg[h][w].p.y = h
			mg[h][w].occupied = false
			mg[h][w].reported = false
			mg[h][w].searched = false
			mg[h][w].lifeValue = rand.Intn(1000) + 1
		}
	}
	return mg
}
