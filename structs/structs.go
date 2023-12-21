package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	person := person{name, age}

	return &person
}

func main() {
    fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	personptr := &person{name: "Ann", age: 40}
	fmt.Println(*personptr)

	fmt.Println(newPerson("Jon",23))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
    fmt.Println((*sp).age)

	(*sp).age = 51
    fmt.Println((*sp).age)

	dog := struct{
		name string
		isGood bool
	}{
		"Jon",
		true,
	}

	fmt.Println(dog)
}
