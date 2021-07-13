package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	var universe Universe
	for h := 0; h < height; h++ {
		universeLine := make([]bool, width)
		universe = append(universe, universeLine)
	}
	return universe
}

func (u Universe) show() {
	for h := 0; h < height; h++ {
		for _, v := range u[h] {
			if v == true {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func (u Universe) seed() {
	nanoseconds := int64(time.Now().Nanosecond())
	rand.Seed(nanoseconds)
	for i := 0; i < width*height/4; i++ {
		h := rand.Intn(15)
		w := rand.Intn(80)
		if u[h][w] == true {
			i--
			continue
		} else {
			u[h][w] = true
		}
	}
}

func (u Universe) alive(h int, w int) bool {
	h, w = getHW(h, w)
	return u[h][w]
}

func (u Universe) neighbors(h int, w int) int {
	h, w = getHW(h, w)
	var neighborNum int
	for i := h - 1; i <= h+1; i++ {
		for j := w - 1; j <= w+1; j++ {
			if i == h && j == w {
				continue
			}
			m, n := getHW(i, j)
			if u[m][n] == true {
				neighborNum++
			}
		}
	}
	return neighborNum
}

func (u Universe) Next(h int, w int) bool {
	h, w = getHW(h, w)
	neighborNum := u.neighbors(h, w)
	var next bool
	if u[h][w] == true {
		if neighborNum < 2 {
			next = false
		} else if neighborNum == 2 || neighborNum == 3 {
			next = true
		} else {
			next = false
		}
	} else {
		if neighborNum == 3 {
			next = true
		} else {
			next = false
		}
	}
	return next
}
