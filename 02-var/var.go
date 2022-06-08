package main

import "fmt"

func main() {
	// 变量是程序运行期间可以改变的量

	var a, b, c int
	fmt.Println("a =", a)
	b = 10
	a = b
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	c = a + b
	fmt.Println("c =", c)
	d := 2333
	fmt.Println(d)

	const e = iota
	fmt.Printf("e = %d\n", e)
}
