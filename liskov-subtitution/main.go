package main

import "fmt"

// Bird interface
type Bird interface {
	Fly() string
}

// Sparrow struct
type Sparrow struct{}

func (s Sparrow) Fly() string {
	return "Sparrow is flying"
}

// Ostrich struct
type Ostrich struct{}

func (o Ostrich) Fly() string {
	return "Ostrich can't fly"
}

// BirdAction takes a Bird and calls its Fly method
func BirdAction(b Bird) {
	fmt.Println(b.Fly())
}

func main() {
	sparrow := Sparrow{}
	ostrich := Ostrich{}

	BirdAction(sparrow)
	BirdAction(ostrich)
}
