package main

import "fmt"

// 递归函数
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return fact(n-1) * n
}

// 闭包递归函数
func main() {

	fmt.Println("fact(5) = ", fact(5))

	// 函数里面定义函数, 必须采用 var 关键字进行声明
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println("fib(5) = ", fib(5))
}
