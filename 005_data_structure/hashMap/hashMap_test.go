package hashmap

import "testing"

func TestPut(t *testing.T) {
	// put int,int
	mp := NewMap()
	for i := 0; i < 100; i++ {
		// 这算不算捕获迭代变量呢？
		// 傻了，这怎么肯算是捕获呢？不会就是发生在 closure/匿名函数/goroutine
		// 这是函数调用
		if err := mp.Put(i, i); err != nil {
			t.Error(err)
		}
	}
	for i := 0; i < 10; i++ {
		if err := mp.Put(i, i); err != nil {
			t.Error(err)
		}
	}

	for i := 0; i < 20; i++ {
		if v, _ := mp.Get(i); v != i {
			t.Fatal("+++++")
		}
	}
}
