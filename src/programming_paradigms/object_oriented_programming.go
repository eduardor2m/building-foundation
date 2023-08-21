package programmingparadigms

import "fmt"

// Methods

type Car struct {
	brand string
	year  int
}

func (c Car) ShowCar() {
	fmt.Println(c.brand, c.year)
}

// Inheritance - Embedded Types (Composition)

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	salary float64
}

func (e Employee) ShowEmployee() {
	fmt.Println(e.name, e.age, e.salary)
}

// Polymorphism

type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() {
	fmt.Println(d.name, "is a dog")
}
