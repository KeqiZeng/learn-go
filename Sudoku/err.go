package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	Errbounds = errors.New("Out of bounds")
	Errdigit1 = errors.New("Invalid digit")
	Errdigit2 = errors.New("Repetitious number in row.")
	Errdigit3 = errors.New("Repetitious number in column.")
	Errdigit4 = errors.New("Repetitious number in subgrid.")
	Errset    = errors.New("Can't set preinstalled digit.")
	Errclear  = errors.New("Can't clear preinstalled digit.")
)

func exitOnError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func inBounds(row, column int) bool {
	if row < 0 || row >= 9 {
		return false
	}
	if column < 0 || column >= 9 {
		return false
	}
	return true
}

func digitValid(digit int) bool {
	return digit > 0 && digit < 10
}
