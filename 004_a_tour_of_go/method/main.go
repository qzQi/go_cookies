package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func useAbser(a Abser) {
	// a.Abs()
	fmt.Println(a.Abs())
}

func printLine() {
	fmt.Println("+++++++++++")
}

func useInterface() {
	// i:=interface{}
	var i interface{}
	describe(i)
	i = 42
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("%v,%T\n", i, i)
}

func useAssert(i interface{}) {
	a, ok := i.(int)
	if ok {
		fmt.Println("is int", a)
	}
	t := i.(string)
	fmt.Println(t)
}

func getType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int type", v)
	case string:
		fmt.Println("string type", v)
	default:
		fmt.Println("other type", v)
	}
}

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}
func usageStringer() {
	p := person{"qi", 21}
	fmt.Println(p)

	var i fmt.Stringer
	i = p
	// s := []int{1, 2, 3}
	// i = s
	fmt.Println(i)

}

type myerror struct {
	when time.Time
	what string
}

// 在需要返回error的时候，只需要返回一个myerror的地址就ok
func (e myerror) Error() string {
	return fmt.Sprintf("at %v,%s happend", e.when, e.what)
}
func runError() error {
	return myerror{time.Now(), "over?"}
}
func usageError() {
	if err := runError(); err != nil {
		fmt.Println(err)
	}
}

func useRead() {
	r := strings.NewReader("hello, reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Println(n, err, b)
		fmt.Println(b[:n])
		if err == io.EOF {
			break
		}
	}
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())
	printLine()
	// a=v error：区分接收者

	useAbser(f)
	useAbser(&v) //&v实现了interface

	printLine()
	useInterface()

	printLine()

	mp := make(map[int]int)
	for i := 0; i < 10; i++ {
		mp[i] = i
	}

	for k := range mp {
		fmt.Println("key is", k)
	}

	// i:=fmt.Println("123")
	var inter interface{}
	inter = mp
	getType(inter)
	inter = "string"
	getType(inter)
	inter = 77
	getType(inter)

	printLine()
	usageStringer()

	printLine()
	usageError()

	printLine()
	useRead()

}
