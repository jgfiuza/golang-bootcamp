// Statement
// Create a function that will receive an interface and it will print the
// interface value with a specific message based on the type.

// Topics to Practice:
// Interfaces, type assertion, switch
package main

import "log"

type structure1 struct {
	value string
}

type structure2 struct {
	value string
}

func (s structure1) GetValue() string {
	return s.value
}

func (s structure2) GetValue() string {
	return s.value
}

type valueInterface interface {
	GetValue() string
}

func printValue(v valueInterface) {
	switch v.(type) {
	case structure1:
		log.Printf("Structure of type 1 with value: %s\n", v.GetValue())
	case structure2:
		log.Printf("Structure of type 2 with value: %s\n", v.GetValue())
	default:
		log.Println("Invalid data type")
	}
}

func main() {
	struc1 := structure1{"First Struct"}
	struc2 := structure2{"Second Struct"}

	printValue(struc1)
	printValue(struc2)
}
