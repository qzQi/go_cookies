package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "now in / process")
	fmt.Println("in /", r.Method, r.URL.String())
	fmt.Fprintln(w, r.URL.String())
	fmt.Fprintln(w, r.URL.Query().Encode())

	// 傻了，傻了。源码自下，了无秘密。不懂的时候打开源码就知道了
	// 其实有了http的基础这些还是很好学的。
	fmt.Println(r.URL.Fragment, r.URL.RawQuery)
}

func handlerQzy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("now in /qzy  process", " ", r.Method, r.URL.String())
	fmt.Fprintln(w, r.URL.String())
	fmt.Fprintln(w, r.URL.Query().Encode())
}

// for post
func handlerPostQzy(w http.ResponseWriter, r *http.Request) {
	// post的数据可以在url也可以在body里面，如何使用？
	// 再学学吧
	// fmt.Println(r.Form.Encode())
	// 其实想想http的东西 就会明白了，request里面有啥东西吗？
	// header？ request line？ body？也就这样嘛
	// r.
	// var a []byte
	// r.Body.Read(a)
	// r.Body.Close()
	// var i interface{}
	// 还是不会用啊、、、问题竟然在json
	// json.Unmarshal(a, &i)
	// mp := i.(map[string]interface{})
	// fmt.Println(mp)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// w.Write() io interface
	obj := map[string]interface{}{
		"name":     "qzy",
		"password": "1234",
	}
	// 向 w里面写数据据就是向w的body里面写
	// w实现了write接口
	// w的定义也很简单就是：WriteHeader  Header
	b, _ := json.Marshal(obj)
	fmt.Fprintf(w, "%s", b)
}

func main() {
	// 如果不增加别的，所有的路由都会由这个处理
	http.HandleFunc("/", handlerRoot)

	// 按照我的感觉这个只能使用/qzy?name=    这样的，不可以使用/qzy/hello
	// 是的就是这样
	http.HandleFunc("/qzy", handlerQzy)

	//现在可以请求，/qzy/hello 页面了
	// 这一套东西的背后的理论知识又是什么？消息的路由吗？
	http.HandleFunc("/qzy/", handlerQzy)

	// for post
	http.HandleFunc("/postqzy", handlerPostQzy)

	// The handler is typically nil, in which case the DefaultServeMux is used.
	// 服务器复用，默认的消息路由
	// 第二个是handler接口
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
