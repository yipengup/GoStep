package main

import "fmt"

// 变长参数
func sum(nums ...int) {
	fmt.Print(nums, "")
	// 将所有的参数进行求和
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// 也可以直接传入一个slice
	s := []int{1, 2, 3, 4}
	// 采用slice...的方式进行传参
	sum(s...)
}
