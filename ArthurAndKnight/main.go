package main

import (
	"fmt"
	"math/rand"
)

type item []string

type character struct {
	name string
	item
}

func (c *character) pickup(i *item) {
	for _, it := range *i {
		c.item = append(c.item, it)
	}
	fmt.Printf("%v pick up something, now he has %v\n", c.name, c.item)
}

func (c *character) give(to *character) {
	i := rand.Intn(len(c.item))
	to.item = append(to.item, c.item[i])
	fmt.Printf("%v give %v to %v\n", c.name, c.item[i], to.name)
	fmt.Printf("Now, %v has %v\n", to.name, to.item)
	c.item[i] = ""
}

func main() {
	arthur := character{name: "Arthur", item: nil}
	knight := character{name: "arm Knight", item: nil}
	var toolbox *item = &item{"hammer", "knife", "sword"}

	fmt.Println("Arthur meet an arm knight, he has nothing.")
	arthur.pickup(toolbox)
	arthur.give(&knight)
}
