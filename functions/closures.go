package main

import "fmt"

// 闭包,匿名函数
// intSeq 返回一个函数
func intSeq() func() int {
	i := 0
	// 这里将 匿名函数返回, 但是隐藏了变量i, 形成闭包
	// 每次调用该函数, 都会将i的值+1
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())

	// 重新调用intSeq函数,返回全新的匿名函数,里面的变量i为新的

	nextInts := intSeq()
	fmt.Println("nextInts:", nextInts())

}
