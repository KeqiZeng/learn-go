package main

import "fmt"

const rows int = 9
const columns int = 9

var subgrid [9]map[int]bool

type number struct {
	value int
	pre   bool
}
type sudoku [9][9]number

func whichsubgrid(row, column int) int {
	var n int
	switch row {
	case 0, 1, 2:
		if column < 3 {
			n = 0
		} else if column >= 3 && column < 6 {
			n = 1
		} else {
			n = 2
		}
	case 3, 4, 5:
		if column < 3 {
			n = 3
		} else if column >= 3 && column < 6 {
			n = 4
		} else {
			n = 5
		}
	case 6, 7, 8:
		if column < 3 {
			n = 6
		} else if column >= 3 && column < 6 {
			n = 7
		} else {
			n = 8
		}
	}
	return n
}

func (s *sudoku) set(row, column int, digit int) {
	var err error
	if s[row][column].pre == true {
		err = Errset
		exitOnError(err)
	}
	if !inBounds(row, column) {
		err = Errbounds
		exitOnError(err)
	}
	if !digitValid(digit) {
		err = Errdigit1
		exitOnError(err)
	}

	for i := 0; i < columns; i++ {
		if digit == s[row][i].value {
			err = Errdigit2
			exitOnError(err)
		}
	}

	for i := 0; i < rows; i++ {
		if digit == s[i][column].value {
			err = Errdigit3
			exitOnError(err)
		}
	}

	n := whichsubgrid(row, column)
	if _, found := subgrid[n][digit]; found {
		err = Errdigit4
		exitOnError(err)
	}

	subgrid[n][digit] = true
	s[row][column].value = digit
}

func (s *sudoku) clear(row, column int) {
	if !inBounds(row, column) {
		err := Errbounds
		exitOnError(err)
	}
	if s[row][column].pre == true {
		err := Errclear
		exitOnError(err)
	}
	n := whichsubgrid(row, column)
	delete(subgrid[n], s[row][column].value)
	s[row][column].value = 0
}

func (s sudoku) show() {
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%2d", s[i][j].value)
		}
		fmt.Printf("\n")
	}
}

func newSudoku(s [rows][columns]int) sudoku {
	var sudo sudoku
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			sudo[i][j].value = s[i][j]
			if s[i][j] != 0 {
				sudo[i][j].pre = true
			}
		}
	}
	for i := 0; i < 9; i++ {
		subgrid[i] = make(map[int]bool, 9)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s[i][j] != 0 {
				subgrid[0][s[i][j]] = true
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 3; j < 6; j++ {
			if s[i][j] != 0 {
				subgrid[1][s[i][j]] = true
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 6; j < 9; j++ {
			if s[i][j] != 0 {
				subgrid[2][s[i][j]] = true
			}
		}
	}
	for i := 3; i < 6; i++ {
		for j := 0; j < 3; j++ {
			if s[i][j] != 0 {
				subgrid[3][s[i][j]] = true
			}
		}
	}
	for i := 3; i < 6; i++ {
		for j := 3; j < 6; j++ {
			if s[i][j] != 0 {
				subgrid[4][s[i][j]] = true
			}
		}
	}
	for i := 3; i < 6; i++ {
		for j := 6; j < 9; j++ {
			if s[i][j] != 0 {
				subgrid[5][s[i][j]] = true
			}
		}
	}
	for i := 6; i < 9; i++ {
		for j := 0; j < 3; j++ {
			if s[i][j] != 0 {
				subgrid[6][s[i][j]] = true
			}
		}
	}
	for i := 6; i < 9; i++ {
		for j := 3; j < 6; j++ {
			if s[i][j] != 0 {
				subgrid[7][s[i][j]] = true
			}
		}
	}
	for i := 6; i < 9; i++ {
		for j := 6; j < 9; j++ {
			if s[i][j] != 0 {
				subgrid[8][s[i][j]] = true
			}
		}
	}
	return sudo
}
