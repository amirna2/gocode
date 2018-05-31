package main

import "fmt"

func main() {
	fmt.Println("UTF-8 runes")

	for i := 900; i < 1000; i++ {
		fmt.Println(string(i), " - ", []byte(string(i)))
	}
}
