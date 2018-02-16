package main

import "fmt"

// Struct definition
type person struct {
	first string
	last  string
	age   int
}

// We can create our own types
type foo int

func main() {
	var myAge foo

	myAge = 15
	p1 := person{"James", "Bond", 31}
	p2 := person{"Miss", "Moneypenny", 25}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Printf("%T %v \n", myAge, myAge)

}
