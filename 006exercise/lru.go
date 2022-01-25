package leetcode

// LRUCache:这个需要map/list啊
// go里面需要自己实现list？
// 自己实现的话，还是使用OO封装一下比较好
// 面试时候自己实现list太麻烦，还是cpp快

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

type list struct {
	head *node
	tail *node
}

func newNode(k, v int) *node {
	return &node{k, v, nil, nil}
}

func newList() *list {
	ans := &list{newNode(-1, -1), newNode(-1, -1)}
	ans.head.next = ans.tail
	ans.head.next.prev = ans.head
	return ans
}

type LRUCache struct {
	cap int
	// lruTable
	lruTable map[int]*node
	lruList  *list
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, make(map[int]*node), newList()}
}

func (lru *LRUCache) Get(key int) int {
	if ptr, ok := lru.lruTable[key]; ok != false {
		value := ptr.value
		delete(lru.lruTable, key)
		tempNode := newNode(key, value)
		ptr.next.prev = ptr.prev
		ptr.prev.next = ptr.next
		// insert tempNode in list head
		tempNode.next = lru.lruList.head.next
		lru.lruList.head.next.prev = tempNode
		tempNode.prev = lru.lruList.head
		lru.lruList.tail = tempNode
		lru.lruTable[key] = tempNode
		return value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	// 如果put超出容量需要进行删除
	tempNode := newNode(key, value)
	if i := len(lru.lruTable); i < lru.cap {
		// 先查看是否存在
		if _, ok := lru.lruTable[key]; ok != false {
			lru.Get(key)
			lru.lruList.head.next.value = value
			return
		}
		// 在链表头部插入节点
		tempNode.next = lru.lruList.head.next
		lru.lruList.tail.prev = tempNode
		tempNode.prev = lru.lruList.head
		lru.lruList.head.next = tempNode
		lru.lruTable[key] = tempNode
	} else {
		// 没有容量了，需要进行删除
		k := lru.lruList.tail.prev.key
		// v:=lru.lruList.tail.prev.value
		delete(lru.lruTable, k)
		ptr := lru.lruTable[k]
		ptr.next.prev = ptr.prev
		ptr.prev.next = ptr.next
		tempNode.next = lru.lruList.head.next
		lru.lruList.tail.prev = tempNode
		tempNode.prev = lru.lruList.head
		lru.lruList.head.next = tempNode
		lru.lruTable[key] = tempNode
	}
}
