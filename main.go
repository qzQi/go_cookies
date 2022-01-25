package main

import (
	"fmt"
	"io"
	"os"
)

type ByteConter int

func (c *ByteConter) Write(p []byte) (int, error) {
	*c += ByteConter(len(p))
	return len(p), nil
}

func main() {
	var c ByteConter
	c.Write([]byte("hello"))
	fmt.Fprintln(&c, "hello world")
	fmt.Println(c)

	var w io.Writer = new(ByteConter)
	// w=ByteConter error:
	// 对接口来说，*ByteConter实现了这个接口
	w = os.Stdout
	// w=new(bytes.Buffer)
	w.Write([]byte("hello world\n"))

	printLine()
	// 只要是从underlying slice赋值而来的，就会改变相同的底层
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[:3]

	// fmt.Printf("%#v \n",s1)
	// fmt.Printf("%#v \n",s2)
	showType(s1, s2)
	fmt.Println("after change s2")
	s2 = append(s2, s2...)
	showType(s2)
	s2[1] = 1
	showType(s1, s2)

	printLine()

	s3 := make([]int, len(s1[:3]))
	copy(s3, s1[:3])
	showType(s1, s3)
	fmt.Println("after change copy")
	s3[len(s3)-1] = 100
	showType(s1, s3)
}
func printLine() {
	fmt.Println("++++++++++++++++++")
}

func showType(elems ...interface{}) {
	// fmt.Printf("%v \n",i)
	for _, v := range elems {
		fmt.Printf("%#v \n", v)
	}
}
