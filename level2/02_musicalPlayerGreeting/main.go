// Musical Player Greetings
// Statement
// Create a Trumpeter struct and a Violinist struct both struts should have
// a Name attribute and a Greetings() function to present themself. Then create
// a MusicalPlayer interface. In the main function, create a slice with two or more
// musical players and for each call the Greetings() function.

// Topics to Practice:
// Type Interfaces, struct, method, function, slice, for each loop
package main

import "fmt"

type Trumpeter struct {
	Name string
}

type Violinist struct {
	Name string
}

func (t Trumpeter) Greetings() string {
	return fmt.Sprintf("Greeting %s, The Trumpeter.", t.Name)
}

func (v Violinist) Greetings() string {
	return fmt.Sprintf("Greeting %s, The Violinist.", v.Name)
}

type MusicalPlayer interface {
	Greetings() string
}

func GreetMusicalPlayer(mp MusicalPlayer) {
	fmt.Println(mp.Greetings())
}

func main() {
	sarah := Violinist{"Sarah Chang"}
	jascha := Violinist{"Jascha Heifetz"}

	bill := Trumpeter{"Bill Adam"}
	maurice := Trumpeter{"Maurice Andre"}

	musicians := []MusicalPlayer{sarah, bill, jascha, maurice}

	for _, musician := range musicians {
		GreetMusicalPlayer(musician)
	}
}
