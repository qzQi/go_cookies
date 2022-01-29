package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Name string
	Body string
	Time int64
	Age  int
}

func main() {
	m := message{"alice", "hello", 12345, 21}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	// b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	fmt.Printf("type is %T ,size %d bytes, toString->%s\n", b, len(b), b)

	// 临时使用的struct用来提取json byte slice里面的特定的内容
	var getName struct {
		Name string
		Age  int
	}

	json.Unmarshal(b, &getName)
	fmt.Println(getName.Name, "age:", getName.Age)

	// 预先不知道byte slice（json字符串）里面的内容
	var f interface{}
	json.Unmarshal(b, &f)
	myMap := f.(map[string]interface{})
	for k, v := range myMap {
		fmt.Printf("k type is %T, v type is %T\n", k, v)
		fmt.Println(k, v)
	}

}
