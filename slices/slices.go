package main

import (
	"fmt"
)

// slice
func main() {

	// 创建slice
	s := make([]string, 3)
	fmt.Println("empty:", s)

	// 使用索引设置或者获取slice的值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// 获取slice的长度
	fmt.Println("len:", len(s))

	// 通过append向slice中添加元素,并返回新的slice对象
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 利用copy函数将原slice中元素拷贝到目标slice中
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 利用slice实现返回包含部分元素的slice
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[2:]
	fmt.Println("sl2:", l)

	l = s[:5]
	fmt.Println("sl3", l)

	// 初始化并声明一个slice
	t := []string{"g", "h", "i"}
	fmt.Println("decl:", t)

	// 二维数组, 并且slice内层的长度是可以改变的
	towD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		// 内部slice长度
		innerLen := i + 1
		// 内部也通过make函数进行声明
		towD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			towD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", towD)

}
