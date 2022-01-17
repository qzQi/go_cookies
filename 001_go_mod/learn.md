# go的module以及package管理
主要的使用：  
* 如何组织代码
* 如何使用别人的
* 如何使用本地的
* 如何使用同一个mod里面不同的package
```go
import (
    "github.com/google/go-cmp/cmp"
    // github.com/google/go-cmp/这是module的名字
    // 但是导入的时候我们需要知道我们所要的package名字
    // 也就是github.com/google/go-cmp/cmp然后go工具会自动分析
    // package所在的位置，然后下载下来
    // 再加一句吧，import导入的package的路径

    "github.com/qzQi/go_cookies/001_go_mod/package_path"
    // 使用这个module里面自己的package的导入方法
)
```

```bash
go mod init github.com/qzQi/go_cookies
# 初始化一个module，module’name可以任意（要发布的话需要通过这个找到）

go mod tidy
# 引用了外部的package后，使用go mod tidy see: go mod

# in go import "example/greetings"
go mod edit -replace example/greetings=../greetings
# （本地module的路径）../greetings是本地的一个module
# 这个用来引用本地的module，感觉不常用
```

哪怕是同一个module里面的文件也不可以通过目录层级来调用。