package main

import (
	"fmt"
	"reflect"
)

func main() {

	t1 := []int{1, 2}
	fmt.Println(reflect.TypeOf(t1))

	t2 := [2]int{1, 2}
	fmt.Printf("%T\n", t2)

	// 初始化数组会被赋予0值
	var a [5]int
	fmt.Println("empty:", a)

	// 利用下标索引进行赋值,获取值
	a[4] = 1
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 获取数组的长度
	fmt.Println("len:", len(a))

	// 声明并初始化数组
	// 这里指定了数组的初始化大小, 不指定的话就会使 slices
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("decl:", b)

	// 二维数组
	var twoD [2][3]int
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[0]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
