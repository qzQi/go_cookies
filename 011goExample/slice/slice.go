package main

import (
	"fmt"
)

/*
func learnSlice() {
	var buffer [256]byte
	// var slice []byte=buffer[100:200]
	// slice=[]byte("aaaaa")
	// a:=[]byte("aaaa")
	slice := buffer[100:200]
	fmt.Println(len(slice))
}

// emmm,确实不熟练了，
func paraSlice(s []int) {
	for i := range s {
		s[i]++
	}
	s = append(s, 1, 2, 3, 4)
	showSlice(s)
}

// 如果想改变slice的长度，必须使用指针或者返回一个slice
func changeLen(slice *[]int) {
	// (*slice).append((*slice),1,2,3)
	// slice.append(slice, 1, 2, 3)
	// 写起来相当麻烦
	(*slice) = append((*slice), 1, 2, 3)
	showSlice(*slice)
}

func showSlice(s []int) {
	fmt.Printf("type is %T \n", s)
	fmt.Println("the size is", len(s), "the cap is", cap(s))
}

func funSlice() {
	// 对slice执行切片结果还是一个slice
	slice := []int{1, 2, 3, 4}
	slice = append(slice, 1, 2, 3)
	// 复制slice的方法，避免修改原slice
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)

	// 扩容原slice
	slice = make([]int, len(slice), 2*cap(slice))

	// copy的des和src可以交织
	// 在pos处插入一个元素
	pos := len(slice) / 2
	slice = slice[0 : len(slice)+1]
	copy(slice[pos+1:], slice[pos:])
	slice[pos] = 1 // the pos to insert

	// 使用append copyslice
	appCopy := append([]int(nil), slice...)
	showSlice(appCopy)
}
func useSnappit() {

}

func cookies() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(len(slice), cap(slice))
	if len(slice) < cap(slice) {
		// access pos:len+1//ok?
		slice = slice[0 : len(slice)+1]
		showSlice(slice)
	} else {
		// len==cap
		// 这个操作当然需要cap> len, have room for new elem
		slice = slice[0 : len(slice)+1]
		showSlice(slice)
	}
}
*/

func seperateLine() {
	fmt.Println("++++++++++++++")
}

func useSlice(s []int) {

}

func buildInFunc() {
	silceInt := []int{1, 2, 3, 4}
	i := len(silceInt)
	// c := cap(silceInt)

	copSlice := make([]int, len(silceInt))
	copy(copSlice, silceInt)

	copSlice = append(copSlice, silceInt...)

	quickCopy := append([]int(nil), copSlice...)
	useSlice(quickCopy)

	// equivalent to the two-line func
	quickCopy = append(make([]int, 0, len(silceInt)), silceInt...)

	// erase i:j
	i, j := 1, 3
	quickCopy = append(quickCopy[:i], quickCopy[j:]...)

}

// erase i--j 左开右闭
func erase(i, j int, s []int) []int {
	s = append(s[:i], s[j:]...)
	return s
}

func printSlice(s []int) {
	fmt.Println(s)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	slice = erase(1, 3, slice)
	printSlice(slice)
}
