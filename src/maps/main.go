package main

import "fmt"

func main() {
	var myGreetings = make(map[string]string)
	myGreetings["Sam"] = "Hello"
	myGreetings["Cece"] = "Bonjour"
	fmt.Println("Sam says  ", myGreetings["Sam"])
	fmt.Println("Cece says ", myGreetings["Cece"])

	myGreetings2 := map[string]string{"Sam": "Hello", "Cece": "Bonjour!"}
	fmt.Println("Cece says ", myGreetings2["Cece"])

	updateMap := func() {
		fmt.Println("---------------------------------------")
		if v, exists := myGreetings2["Dad"]; exists {
			fmt.Println("Dad greetings exists \"", v, "\" removed it ", myGreetings2)
			delete(myGreetings2, "Dad")
			fmt.Println("Removed Dad -> ", myGreetings2)
		} else {
			myGreetings2["Dad"] = "Salut!"
			fmt.Println("Dad greetings does not exist. Added it: ", myGreetings2)
		}
		for k, val := range myGreetings2 {
			fmt.Println(k, "says ", val)
		}
	}
	updateMap()
	updateMap()

}
