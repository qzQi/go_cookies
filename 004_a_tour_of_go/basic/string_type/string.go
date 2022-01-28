package main

import (
	"fmt"
)

func forLoop(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
}

func forRange(str string) {
	for i, v := range str {
		fmt.Println(i, "  ", v)
	}
}

func main() {
	name := "祁志赟"
	fmt.Println(len(name))
	for i, v := range name {
		fmt.Printf("at pos %d,is %v\n", i, v)
	}
	// 对非 ASCII string使用index无意义，使用for range才能解释
	fmt.Println(name[1])

	sample := "��=� ⌘"
	forLoop(sample)
	forRange(sample)
}
