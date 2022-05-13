package main

import "fmt"

// 定义一个值传递函数
// 值传递实际上是原始数据的一个副本, 对形参的修改不会应到到原始数据
func zeroval(value int) {
	value = 0
}

// 定义一个参数为 int类型指针的函数
func zeroptr(iptr *int) {
	// 通过 *iptr 获取到对应地址的值,此时会修改指针对应的原始数据
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// 通过 &i 获取到对应的指针
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 打印对应的指针信息
	fmt.Println("pointer:", &i)
}
