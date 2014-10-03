package main

import "fmt"

type People struct {
	name string
}

// COMPOSITION START
type Employee struct {
	People    // HL
	emloyeeID int
}

type Driver struct {
	People       // HL
	licensePlate int
}

//COMPOSITION END OMIT

func (p People) Greet() string {
	return fmt.Sprintf("Hello, my name is %s ", p.name)
}

// FIELDS START OMIT
func (e Employee) Greet() string {
	return fmt.Sprintf("Hello, my name is %s (employeenumber %d)  ", e.name, e.emloyeeID) // HL

}

// FIELDS END OMIT

func (d Driver) Greet() string {
	return fmt.Sprintf("Hello, my name is %s (my plate is %d)", d.name, d.licensePlate)
}

// Greet END OMIT

// main START OMIT
func main() {
	p := People{"Finch"}
	e := Employee{p, 1}
	d := Driver{p, 2}
	fmt.Println(p.Greet())
	fmt.Println(e.Greet())
	fmt.Println(d.Greet())

}

// main END OMIT

// main END OMIT
