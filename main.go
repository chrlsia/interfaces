package main

import "fmt"

//Animal interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

// Dog is a type for dogs
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

//Says is a method of Dog type
func (d Dog) Says() string {
	return d.Sound
}

//HowManyLegs is a method of Dog type
func (d Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

// Cat is a type for cats
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

//Says is a method of Cat type
func (c Cat) Says() string {
	return c.Sound
}

func (c Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main() {
	//ask a riddle = γρίφος, σπαζοκεφαλιά
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	Riddle(dog)

	var cat Cat
	cat.Name = "cat"
	cat.Sound = "meow"
	cat.NumberOfLegs = 4
	cat.HasTail = true

	Riddle(cat)
}

//Riddle gets a parameter of type Animal
func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal "says" %s and has %d legs.`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}

// No need to have pointer , cause nothing is going to be changed to
// the dog and cat variable
