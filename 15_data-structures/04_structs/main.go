package main

import "fmt"

// Struct definition
type Person struct {
	first string
	last  string
	age   int
}

type DoubleZero struct {
	//anonymous type. DoubleZero has directly access to Person fields
	Person
	First         string
	LicenseToKill bool
}

// We can create our own types
type foo int

func main() {
	var myAge foo

	myAge = 15

	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   31},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}
	p2 := Person{"Miss", "Moneypenny", 25}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Printf("%T %v \n", myAge, myAge)

}
