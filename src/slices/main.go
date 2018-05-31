package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 4)
	fmt.Println("-------------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	for i := 0; i < 20; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), " Cap:", cap(mySlice), " Value:", mySlice[i])
	}
	fmt.Println("-------------------------")

	mySlice2 := []string{
		"Zero",
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
	}
	fmt.Println("[1:4] ", mySlice2[1:4])
	fmt.Println("[:4]  ", mySlice2[:4])
	fmt.Println("[:]   ", mySlice2[:])
	fmt.Println("[4:]  ", mySlice2[4:])

	//var mySlice3 = []string{}
	mySlice3 := make([]string, 0, len(mySlice2))
	//build a new slice taking only even elements
	for i := 0; i < len(mySlice2); i++ {
		if i%2 == 0 {
			mySlice3 = append(mySlice3, mySlice2[i])
		}
	}
	fmt.Println(mySlice3)
	// Leaving the "Two" element out
	fmt.Println(append(mySlice3[:1], mySlice3[2:]...))
}
