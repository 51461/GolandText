package main

import "fmt"

func add(i, j int) (a int) {
	return i + j
}
func main() {
	fmt.Println("hello world")
	var max = maxfunc(5, 10)
	var x int = add(1, 2)
	fmt.Println("最大值", max)
	fmt.Println("和", x)

}

//求两数最大值
func maxfunc(a, b int) (c int) {
	if a > b {
		c = a
	} else {
		c = b
	}
	return c
}
