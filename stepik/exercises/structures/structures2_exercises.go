package main

import "fmt"

type Person struct {
	name string
}
type Android struct {
	Person
	model string
}

func (android *Android) talk() {
	android.Person.talk()
	fmt.Println("But not really human... And not really " + android.name)
}

func (person *Person) talk() {
	fmt.Println("HI I AM HUMAN. MY NAME IS " + person.name + ".")
}

func main() {
	var android = Android{
		model: "model",
		Person: Person{
			name: "PAVLO",
		},
	}

	android.talk()
}
