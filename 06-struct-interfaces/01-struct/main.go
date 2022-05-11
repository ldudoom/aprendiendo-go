package main

import "fmt"

// Struct Person

type Person struct {
	name string
	age  int
}

func (p *Person) toString() string {
	str := fmt.Sprintln("La persona", p.name, "tiene la edad de", p.age, "anios")
	return str
}

func main() {
	p1 := Person{"Raul", 41}
	p1.name = "Alejandro"
	p1.age = 44
	fmt.Print(p1.toString())

	p2 := Person{
		name: "Fernanda",
		age:  40,
	}
	fmt.Println(p2.toString())
}
