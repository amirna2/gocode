package main

import "fmt"

// simple function
func greet(fname, lname string) {
	fmt.Printf("Hello %v\n", fmt.Sprint(fname, " ", lname))
}

// fuction with multiple returns
func greet2(fname, lname string) (string, string) {
	return "Hi", fmt.Sprint(fname, " ", lname)
}

// variadic function
func greet3(name ...string) {
	fmt.Printf("Hello ")
	for _, s := range name {
		fmt.Printf("%v ", s)
	}
	fmt.Printf("\n")

}

// function as argument (callback)
func greet4(callback func(name string), names ...string) {
	fmt.Printf("Hello")
	for _, n := range names {
		callback(n)
	}
	fmt.Printf("\n")
}

func main() {

	var s1, s2 string

	// defered the execution of this function until all functions below have executed and we are about to exit
	// the main function
	defer greet("Amir", "Nathoo")

	s1, s2 = greet2("Amir", "Nathoo")
	fmt.Printf("%v %v\n", s1, s2)

	greet3("Amir", "Emily", "Cece", "Sam")
	//slice (or array) of strings
	names := []string{"Amir", "Emily", "Cece", "Sam"}
	greet3(names...)

	//anonymous function
	s := func(fname, lname string) (string, string) {
		return "Hi", fmt.Sprint(fname, " ", lname)
	}

	s1, s2 = s("Amir", "Nathoo")
	fmt.Printf("%v %v\n", s1, s2)

	greet4(
		func(name string) {
			fmt.Printf(" %v", name)
		},
		names...)
}
