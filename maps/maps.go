package main

import "fmt"

// map结构
func main() {

	// 使用 make + map 函数创建空的 map
	m := make(map[string]int)
	fmt.Println(m)

	// 使用类似于数组去设置值
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)

	// 使用类似于数组去获取值
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// 不存在的数据 会显示对应的零值, 当前int就是0
	v3 := m["k3"]
	fmt.Println("v3:", v3)
	fmt.Printf("%T\n", v3)

	// len 函数显示键值对的个数
	fmt.Println("len:", len(m))

	// 通过delete 函数删除map中对应的主键
	delete(m, "k1")
	fmt.Println("m:", m)

	// 可以使用 _ 表示不会使用这个参数
	// 第二个参数表示 map里面是否存在指定的键
	_, prs := m["k4"]
	fmt.Println("prs:", prs)

	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("m2", m2)
}
