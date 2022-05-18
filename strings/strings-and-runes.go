package main

import (
	"fmt"
	"unicode/utf8"
)

// 字符串和runes区别
// rune 字符, 表示unicode编码的整数
func main() {

	// 在go中 字符串相当于只读的 byte 切片
	// go的字符串是utf8编码的
	const s = "你好"

	// 因为字符串底层相当于 []byte , 获取到底层字节数组的长度
	fmt.Println("Len:", len(s))

	// 获取到底层byte数组的每一个byte的unicode码
	for i := 0; i < len(s); i++ {
		// 以16进制输出unicode码
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("")

	// 利用utf8包计算字符串中包含多少个rune,也就是多少个utf8编码的字符
	fmt.Println("RuneCount:", utf8.RuneCountInString(s))

	// 利用range 获取到的是每个rune字符在 底层byte数组的中的偏移量, 这里只会循环两次
	for index, runeValue := range s {
		fmt.Printf("%U starts at %d\n", runeValue, index)
	}

	fmt.Println("Using DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		// 利用 DecodeRuneInString 获取到当前字符串byte数组中表示一位rune 以及其长度
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%U starts at %d\n", runeValue, i)
		w = width
	}
}
