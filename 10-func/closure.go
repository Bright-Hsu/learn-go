package main

import "fmt"

func main() {
	addFunc := myAdd(1, 2)
	fmt.Println(addFunc())
	fmt.Println(addFunc())
	fmt.Println(addFunc())
}

func myAdd(x1, x2 int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x1 + x2
	}
}
