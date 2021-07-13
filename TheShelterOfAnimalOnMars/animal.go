package main

import (
	"fmt"
)

type Mover interface {
	move()
}

type Eater interface {
	eat()
}

type Sleeper interface {
	sleep()
}

func AnimalMove(m Mover) {
	m.move()
}

func AnimalEat(e Eater) {
	e.eat()
}

func AnimalSleep(s Sleeper) {
	s.sleep()
}

type Dog struct {
	name string
	food string
}

func (d Dog) move() {
	fmt.Printf("Dog %v is moving to somewhere\n", d.name)
}

func (d Dog) eat() {
	fmt.Printf("Dog %v is eating %v\n", d.name, d.food)
}

func (d Dog) sleep() {
	fmt.Printf("Dog %v is sleeping\n", d.name)
}

func (d Dog) String() string {
	return d.name
}

type Cat struct {
	name string
	food string
}

func (c Cat) move() {
	fmt.Printf("Cat %v is moving to somewhere\n", c.name)
}

func (c Cat) eat() {
	fmt.Printf("Cat %v is eating %v\n", c.name, c.food)
}

func (c Cat) sleep() {
	fmt.Printf("Cat %v is sleeping\n", c.name)
}

func (c Cat) String() string {
	return c.name
}

type Bear struct {
	name string
	food string
}

func (b Bear) move() {
	fmt.Printf("Bear %v is moving to somewhere\n", b.name)
}

func (b Bear) eat() {
	fmt.Printf("Bear %v is eating %v\n", b.name, b.food)
}

func (b Bear) sleep() {
	fmt.Printf("Bear %v is sleeping\n", b.name)
}

func (b Bear) String() string {
	return b.name
}
