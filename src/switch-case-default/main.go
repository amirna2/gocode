package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Printf("%v is of type int\n", x)
	case string:
		fmt.Printf("%v is of type string\n", x)
	case Person:
		fmt.Printf("%v is of type Person\n", x)
	default:
		fmt.Printf("%v is of type unknown\n", x)
	}
}

func (p *Person) changeFirstName(fn string) {
	p.firstName = fn
}
func main() {
	name := "Max"

	switch name {
	case "Max":
		fmt.Print("Hello Max and")
		fallthrough
	case "Sam":
		fmt.Println(" Sam")

	case "Amir":
		greetings := "Salut"
		fmt.Println(greetings, " ", name)
	default:
		fmt.Println("Hello!")
	}
	fmt.Println("========================================")

	switchOnType(7)
	var t = Person{"Amir", "Nathoo"}
	t.changeFirstName("Joe")
	fmt.Println(t)
	switchOnType("Amir")
	switchOnType(10.5)
	switchOnType(t)

}
