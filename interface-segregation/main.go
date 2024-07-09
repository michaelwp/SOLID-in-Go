package main

import "fmt"

// Worker interface with specific methods
type Worker interface {
	Work()
}

// Eater interface with specific methods
type Eater interface {
	Eat()
}

// Human struct implementing both Worker and Eater
type Human struct{}

func (h Human) Work() {
	fmt.Println("Human is working")
}

func (h Human) Eat() {
	fmt.Println("Human is eating")
}

// Robot struct implementing only Worker
type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robot is working")
}

func main() {
	var worker Worker

	worker = Human{}
	worker.Work()

	worker = Robot{}
	worker.Work()

	var eater Eater = Human{}
	eater.Eat()
}