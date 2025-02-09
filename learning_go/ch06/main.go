package main

import (
	"errors"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func UpdateSlice(strings []string, s string) error {
	length := len(strings)
	if length == 0 {
		return errors.New("Empty slice")
	}
	strings[length-1] = s
	fmt.Println(strings)
	return nil
}

func GrowSlice(strings []string, s string) {
	strings = append(strings, s)
	fmt.Println(strings)
}

func main() {

	// ex1
	person := MakePerson("Monica", "Geller", 26)
	fmt.Println(person)
	personPointer := MakePersonPointer("Rachel", "Green", 26)
	fmt.Println(personPointer)

	// ex2
	s := "Joey"
	strings := []string{"Monica", "Rachel"}
	UpdateSlice(strings, s)
	GrowSlice(strings, s)
	fmt.Println(strings)

	// ex3
	//var people []Person
	people := make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		people = append(people, MakePerson("Rachel", "Green", 26))
	}

}
