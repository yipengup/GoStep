package main

import "fmt"

// 函数多个返回值, 返回类型用括号进包裹
func vals() (int, int) {
	return 3, 7
}

func main() {
	i, i2 := vals()
	fmt.Println("i=", i)
	fmt.Println("i2=", i2)

	// 当不需要某个返回值时,使用 _ 即可
	_, i3 := vals()
	fmt.Println("i3=", i3)
}
