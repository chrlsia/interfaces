package main

import "fmt"

// interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

// Dog is a type for dogs
type Dog struct{
	Name string
	Sound string
	NumberOfLegs int
}

func (d Dog) Says() string {
	return d.Sound
}

func (d Dog) HowManyLegs() int {
	return d.NumberOfLegs
}
// Cat is a type for cats
type Cat struct {
	Name string
	Sound string
	NumberOfLegs int
	HasTail bool
}

func (c Cat) Says() string {
	return c.Sound
}

func (c Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main(){
	//ask a riddle = γρίφος, σπαζοκεφαλιά
	dog:=Dog{
		Name: "dog",
		Sound:"woof",
		NumberOfLegs: 4,
	}

	Riddle(dog)

	var cat Cat
	cat.Name ="cat"
	cat.Sound ="meow"
	cat.NumberOfLegs = 4
	cat.HasTail = true

	Riddle(cat)
}

func Riddle(a Animal) {
	riddle:= fmt.Sprintf(`This animal "says" %s and has %d legs.`, a.Says(),a.HowManyLegs())
	fmt.Println(riddle)
}

// No need to have pointer , cause nothing is going to be changed to
// the dog and cat variable