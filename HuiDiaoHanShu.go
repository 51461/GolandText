package main

import "fmt"

/**
回调函数demo
	回调函数。函数参数是函数类型，这个函数就是回调函数
*/
func main() {
	var result int
	result = funccc(20, 10, jia)
	fmt.Println(result)
	result = funccc(20, 10, jian)
	fmt.Println(result)
	result = funccc(20, 10, cheng)
	fmt.Println(result)
	result = funccc(20, 10, chu)
	fmt.Println(result)
}

type funcType func(int, int) int

func funccc(a, b int, fun funcType) (c int) {
	fmt.Print("funccc:")
	c = fun(a, b)
	return
}
func jia(a int, b int) int {
	return a + b
}
func jian(a int, b int) int {
	return a - b
}
func cheng(a int, b int) int {
	return a * b
}
func chu(a int, b int) int {
	return a / b
}
