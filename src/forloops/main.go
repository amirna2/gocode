package main

import "fmt"

func main() {
	fmt.Printf("For loop 1: iterration\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	fmt.Printf("For loop 2: conditional. Same as while(condition){} in C\n")
	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Printf("\n")

	fmt.Printf("For loop 3:  infinite. Same as while (true) {} in C\n")
	i = 0
	for {
		fmt.Printf("%d ", i)
		i++
		if i > 9 {
			break
		}
	}

	fmt.Printf("\n")

	fmt.Printf("For loop 4:  continue if even number, otherise print\n")
	i = 0
	for {
		i++
		if i%2 == 0 {
			continue
		} else {
			fmt.Printf("%d ", i)
		}
		if i > 20 {
			break
		}
	}
}
