/*
hashmap:
实现这个结构需要考虑的事情：
1、hash算法 hash(key)-->where to place
2、hash conflict：解决冲突的策略
*/
package hashmap

import (
	"errors"
	"fmt"
	"hash/fnv"
)

type node struct {
	key   interface{}
	value interface{}
	next  *node
}

type HashMap struct {
	capacity uint64
	size     uint64
	table    []*node
}

var defaultSize uint64 = 1 << 10

func NewMap() *HashMap {
	return &HashMap{
		capacity: defaultSize,
		size:     0,
		table:    make([]*node, defaultSize),
	}
}

func newNode(key, value interface{}) *node {
	return &node{key, value, nil}
}

func (hm *HashMap) hash(key interface{}) uint64 {
	h := fnv.New64a()

	h.Write([]byte(fmt.Sprintf("%v", key)))

	hashVal := h.Sum64()
	return (hm.capacity - 1) & (hashVal ^ (hashVal >> 16))
}

// Put如果已经存在返回err
func (hm *HashMap) Put(k, v interface{}) error {
	pos := hm.table[hm.hash(k)]
	n := newNode(k, v)
	// pos=append(pos,n)
	for cur := pos; cur != nil; cur = cur.next {
		if cur.key == k {
			return fmt.Errorf("already exist k:%v  v:%v\n", k, v)
		}
	}
	n.next = pos
	hm.table[hm.hash(k)] = n
	hm.size++
	return nil
}

func (hm *HashMap) Get(k interface{}) (interface{}, error) {
	head := hm.table[hm.hash(k)]
	if head == nil {
		return nil, errors.New("no value")
	}

	for cur := head; cur != nil; cur = cur.next {
		if cur.key == k {
			return cur.value, nil
		}
	}
	return nil, errors.New("not found")
}

func (hm *HashMap) Size() uint64 {
	return hm.size
}
