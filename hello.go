package main

import "fmt"

type Speaker interface {
	Speak() error
}

type Dog struct{}

func (d *Dog) Speak() error {
	fmt.Println("BowWow")
	return nil
}

type Cat struct{}

func (c *Cat) Speak() error {
	fmt.Println("Meow")
	return nil
}

func DoSpeak(s Speaker) error {
	return s.Speak()
}

func main() {
	dog := Dog{}
	DoSpeak(&dog) // "BowWow"が表示される

	cat := Cat{}
	DoSpeak(&cat) // "Meow"が表示される
}
