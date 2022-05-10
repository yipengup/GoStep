package main

import "fmt"

// range 用于迭代各种数据结构
func main() {

	s := []int{1, 2, 3}
	// for range 循环, 第一个参数为索引, 第二个参数为值
	sum := 0
	for _, num := range s {
		sum += num
	}
	fmt.Println("sum", sum)

	for index, num := range s {
		if num == 3 {
			fmt.Println("index:", index)
		}
	}

	// 用range循环迭代map
	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 仅仅获取map的key值
	for key := range m {
		fmt.Println("key:", key)
	}

	// range 循环迭代字符串
	// 第一个参数是字符在字符串所在的位置索引, 第二个参数是字符对应的unicode码
	for i, c := range "go冲" {
		fmt.Println(i, c)
	}

}
