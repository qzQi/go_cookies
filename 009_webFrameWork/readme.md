go moudle还是不太会组织，这个在内部创建的package，可能目前外部不需要
引用，可以在内部init一个 mod？

This file is in D:\vscode\go_cookies\009_webFrameWork\gee\day01, which is a nested module in the D:\vscode\go_cookies module.
gopls currently requires one module per workspace folder.
Please open D:\vscode\go_cookies\009_webFrameWork\gee\day01 as a separate workspace folder.
You can learn more here: https://github.com/golang/tools/blob/master/gopls/doc/workspace.md

一个工作区域（mod）里面不能内嵌mod。


### 01
```go
package http

type Handler interface{
    ServeHttp(w http.ResponseWriter,r *http.Request)
}
//  The handler is typically nil, in which case the DefaultServeMux is used
func ListtenAndServe(address string,h Handler)
```
这里解释我们使用内置的http，最后传入一个nil就是使用默认的 serveMux。

也就是说只要我们实现了这个接口就可以接管所有的http请求了。


关于http.ReaponseWriter && http.Request.

web服务无非就是根据request构建出我们的writer。   
request里面有哪些东西？response我们需要填入哪些东西？

这都是我们http协议里面的知识，大家都不陌生。      
for request： requestLine/header/body

for response: responseLine/header/body       
responseWriter需要先set然后write Header。
```go
	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //在set之后，先set后write
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
```