package main

import "fmt"

type Person struct {
	name string
}
type Android struct {
	person Person
	model  string
}

func (person *Person) talk() {
	fmt.Println("HI I AM HUMAN")
}

func main() {
	var android = Android{
		model: "model",
		person: Person{
			name: "PAVLO",
		},
	}

	android.person.talk()
}
