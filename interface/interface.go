package interfaces

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct { name string }
type Cat struct { name string }

func (d Dog) Speak() {
	fmt.Println(d.name, "speak owhh!")
}

func (c Cat) Speak() {
	fmt.Println(c.name, "speak meow!")
}

func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

func InterfaceExample() {
	dog := Dog{name: "Baymax"}
	dog.Speak()

	cat := Cat{name: "pinky"}
	cat.Speak()
}