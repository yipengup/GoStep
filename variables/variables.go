package main

import "fmt"

/* 声明变量 */
func main() {

	// 声明一个变量，编译器会自动进行类型推断
	var a = "initial"
	fmt.Println(a)

	// 同时声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// 没有进行初始化复制的变量被赋予0值
	var e int
	fmt.Println(e)

	// 通过 := 语法快速声明变量以及为变量初始化
	f := "apple"
	fmt.Println(f)

}
