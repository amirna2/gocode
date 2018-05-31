package main

import "fmt"

// declared package scope - zero value
var x float64

const q = 42
const r string = "A constant string"

const (
	I = iota // 0
	J = iota // 1
	K = iota // 2
)

// declared package scope - pre initialized
var e, f int = 10, 100

func main() {
	// short hand
	a := true
	b := 10.5
	c := "Hello, 世界"
	d := 51
	fmt.Printf("%v (%T)\n", a, a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v (%T)\n", x, x)
	fmt.Printf("constant q=%v (%T)\n", q, q)
	fmt.Printf("constant r=\"%v\" (%T)\n", r, r)
	fmt.Printf("constant I=%v J=%v K=%v\n", I, J, K)

}
