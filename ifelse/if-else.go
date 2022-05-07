package main

import (
	"fmt"
)

// 在 go 里面，条件可以不用括号，但是大括号是必须的
func main() {

	// 基本用法
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//  只有if语句
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 条件分支前面可以声明变量，变量可以在分支里面使用
	if number := 9; number < 0 {
		fmt.Println(number, "is negative")
	} else if number < 10 {
		fmt.Println(number, "has 1 digit")
	} else {
		fmt.Println(number, "has multiple digits")
	}

}
