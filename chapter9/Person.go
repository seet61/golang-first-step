package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	a := Android{Person{"LG"},
		"7.1.1.1"}
	//a.Person.Talk()
	a.Talk()
	fmt.Println("My model:", a.Model)
}
