发现go的blog真是一个好东西。

https://go.dev/blog/json
介绍go json。

能够展示为合法的json形式的数据结构都可以被编码为json文件。

关于json的具体可以看我写的json paser for C++


json数组与object都可以作为json文件。


结构体的话，只有可导出的才可以被解析？

The json package only accesses the exported fields of struct types (those that begin with an uppercase letter). Therefore only the the exported fields of a struct will be present in the JSON output.

看那篇blog里面有你关于go解析json的一切。

使用json数据填populate充我们的数据结构。

用的时候再来看一下这个blog吧