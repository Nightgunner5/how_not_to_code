package main

import "fmt"

func main() {
	var a bool
	var b bool
	var c, d, e bool
	var f bool

	a = !b
	c = c != a || b
	d = b || a && c
	a = e == d == e
	f = !a
	a = f

	if !a != false == true {
		fmt.Print("You can see ")
	}
	if b != false && true {
		fmt.Print("anchovy pizza ")
	}
	if true == c && !true == false {
		fmt.Print("why this would be ")
	}
	if (false && true || false && false != true || true == false && true) == d {
		fmt.Print("edible ")
	}
	if (e || f) == e {
		fmt.Println("confusing.")
	}
}
