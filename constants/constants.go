package main

import (
	"fmt"
	"math"
)

// go 支持 字符、字符串、boolean、数字型的常量

// 声明字符串常量
const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 50000000

	// 3e20 = 3 * 10^20
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	// 数字型常量会根据上下文转换为其他需要类型， 这里的n会被转换成 float
	fmt.Println(math.Sin(n))

}
