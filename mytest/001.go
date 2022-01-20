// package main

// import (
// 	lib001 "github.com/qzQi/go_cookies/001_go_mod"
//  如果路径的最后与package name不一样的话起别名？
// )

// func main() {
// 	lib001.Add(1, 2)
// 	// 测试文件显然不是可见的
// 	// 这里是自己编写的测试用例/以及使用
// 	// package内的测试就在package内部
// }
package main

import (
	lib001 "github.com/qzQi/go_cookies/001_go_mod"
)

func main() {
	lib001.Add(1, 2)
}
