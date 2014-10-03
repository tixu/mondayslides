package main

import "fmt"

// People START OMIT
type People struct {
	name string
}

// People END OMIT

// Greet START OMIT
func (p People) Greet() string {
	return fmt.Sprintf("Hello, my name is %s ", p.name)
}

// Greet END OMIT

// main START OMIT
func main() {
	p := People{"Finch"}
	fmt.Println(p.Greet())
	fmt.Println(People{"Reese"}.Greet())

}

// main END OMIT
