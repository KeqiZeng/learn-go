package main

import (
	"fmt"
	"time"
)

func main() {
	universe1 := NewUniverse()
	universe2 := NewUniverse()
	universe1.seed()
	fmt.Printf("\nThe 1st generation\n")
	universe1.show()
	for i := 0; i < 9; i++ {
		universe2 = Step(universe1)
		universe1, universe2 = universe2, universe1
		switch i + 2 {
		case 2:
			fmt.Printf("\nThe %vnd generation\n", i+2)
		case 3:
			fmt.Printf("\nThe %vrd generation\n", i+2)
		default:
			fmt.Printf("\nThe %vth generation\n", i+2)
		}
		universe1.show()
		time.Sleep(5)
	}
}
