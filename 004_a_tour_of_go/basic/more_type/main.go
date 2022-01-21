package main

import (
	"fmt"
	// "go/types"
)

// Vertex is x y
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Printf("%T\n", v)
	p := &v
	(*p).X = 1e9
	p.X = 1
	fmt.Println(v)
	sliceTest()
	printLine()

	var s1 []int = []int{1, 2, 3, 45}
	// var s2 =[]int{1,2,3,4}
	// s3:=[]int{1,2,3,4}
	printSlice(s1)
	// s2 := s1[1:]
	// s1 = append(s1, 1, 2, 3, 4)
	printSlice(s1)
	for len(s1) == cap(s1) {
		s1 = append(s1, 1, 2)
		printSlice(s1)
	}

	if s1 != nil {
		fmt.Println("slice could cmp with nil")
	}
	printLine()

	s2 := make([]int, 5)
	s2 = append(s2, 1, 2, 3)
	printSlice(s2)
	len, cap := 5, 5
	s3 := make([]int, len, cap)
	printSlice(s3)

	printLine()
	fn := closure2()
	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}

}

func sliceTest() {
	q := []int{2, 3, 45, 7}
	fmt.Printf("%T\t %q\n", q, q)

	r := []bool{false, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{1, false},
		{2, true},
		{3, false},
	}
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d %#v\n", len(s), cap(s), s)
}

func printLine() {
	fmt.Println("++++++++++++++")
}

func usageMap() {
	// init
	var m map[string]int = map[string]int{"qi": 1, "zhi": 2}
	fmt.Println(m)
	m2 := m
	m3 := make(map[string]int)
	m3 = m
	fmt.Println(m, m2, m3)
}

// go如何使用闭包呢？，也没有function/bind/lambda这样的
// 神兵利器
func usageFunc(fn func(int, int) int) int {
	return fn(3, 4)
}

func closure() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func closure2() func() int {
	s := []int{1, 2, 3, 4}
	i := 1
	return func() int {
		i++
		s = append(s, i)
		fmt.Println(s)
		return s[len(s)-1]
	}
}
