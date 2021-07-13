package main

import (
	"fmt"
)

func main() {
	s := newSudoku([rows][columns]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
	s.set(0, 2, 1)
	s.clear(0, 2)
	s.set(1, 1, 4)
	s.set(7, 7, 3)
	s.show()
	fmt.Println(subgrid[0])
	fmt.Println(subgrid[8])
}
