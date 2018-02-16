package main

import "fmt"

// Struct definition
type Person struct {
	First string
	Last  string
	Age   int
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

	fmt.Println(p1.First, p1.Last, p1.Age)
	fmt.Println(p2.First, p2.Last, p2.Age)
	fmt.Printf("%T %v \n", myAge, myAge)

}
