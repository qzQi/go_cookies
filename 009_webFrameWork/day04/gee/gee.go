package gee

import "net/http"

// "os"
// "log"

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine //all group have the same engine
}

type Engine struct {
	*RouterGroup //在OO里面这算inherit？go里面算是组合？通过Engine来调用RouterGroup的方法
	//   只不过是一个语法糖，编译器自动的展开了
	// eg: e.NewGroup()=> e.RouterGroup.NewGroup()
	router *router
	groups []*RouterGroup //store all groups
}

func NewEngine() *Engine {
	e := &Engine{}
	e.router = newRouter()
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup}
	return e
}

// NewGroup share the same underlying Engine instance
// init..ly this will bw called by Engine
// e:=NewEngine();  e.NewGroup("/qzy");
func (g *RouterGroup) NewGroup(prefix string) *RouterGroup {
	engine := g.engine
	ans := &RouterGroup{}
	ans.engine = engine
	ans.parent = g
	ans.prefix = g.prefix + prefix

	engine.groups = append(engine.groups, ans)
	return ans
}

func (g *RouterGroup) addRouter(method, pattern string, handler HandlerFunc) {
	pattern = g.prefix + pattern
	// log.Println()

	g.engine.router.addRoute(method, pattern, handler)

}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRouter("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRouter("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	contex := newContex(w, r)
	e.router.handle(contex)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func test() {
	e := NewEngine()
	e.NewGroup("/api")
}
