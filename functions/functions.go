package main

import "fmt"

// 两个int参数, 返回值为int
// go必须明确返回值
func plus(a int, b int) int {
	return a + b
}

// 相同参数类型可以合并一起声明
func plusPlus(a, b, c int) int {
	return a + b + c
}

// 函数
func main() {

	i := plus(1, 2)
	fmt.Println("1+2=", i)

	fmt.Println("1+2+3=", plusPlus(1, 2, 3))

}
