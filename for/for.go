package main

import "fmt"

// for 循环是 go里面唯一支持的循环结构
func main() {
	i := 1
	// 单独条件
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for循环的基本用法
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	// 无条件 类似于white true 循环
	for {
		fmt.Println("loop")
		// 通过 break 或者 return 结束循环
		break
	}

	// 通过continue结束当前循环，执行下次循环
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
