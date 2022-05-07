package main

import (
	"fmt"
	"time"
)

// switch 语句
func main() {
	i := 2
	fmt.Print("Write ", i, " as ")

	// 基本的switch 语句
	switch i {
	case 1:
		fmt.Println("one")
		// go语言中不需要使用 break 关键字
		// break
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	// 用逗号分割多个相同的case
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	fmt.Println("t", t)

	// 无条件的switch是if-else的另一种形式
	switch {
	case t.Hour() < 12:
		fmt.Println("It's a before noon")
	default:
		fmt.Println("It's a after noon")
	}

	// 定义一个函数, interface{} 表示一个空接口类型, 可以类比于Java中的Object, 表明可以传递任意的参数类型
	whatAmI := func(i interface{}) {
		// 这里直接获取到 参数的类型
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			// printf 可以对字符串进行格式化
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
