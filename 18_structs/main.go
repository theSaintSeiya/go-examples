package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 23
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(&person{"Bob", 20})
	fmt.Println(newPerson("Tom"))

	fmt.Println(person{name: "Fred"})

	s := person{name: "Tom"}
	sp := &s

	sp.age = 5
}
