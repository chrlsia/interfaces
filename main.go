package main

import "fmt"

// interface type

// Dog is a type for dogs
type Dog struct{
	Name string
	Sound string
	NumberOfLegs int
}

// Cat is a type for cats
type Cat struct {
	Name string
	Sound string
	NumberOfLegs int
	HasTail bool
}

func main(){
	//ask a riddle = γρίφος, σπαζοκεφαλιά
	dog:=Dog{
		Name: "dog",
		Sound:"woof",
		NumberOfLegs: 4,
	}

	Riddle(dog)
}

func Riddle(d Dog) {
	riddle:= fmt.Sprintf(`This animal "says" %s and has %d legs.`, d.Sound,d.NumberOfLegs)
	fmt.Println(riddle)
}