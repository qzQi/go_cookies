package list

type node struct {
	val  interface{}
	prev *node
	next *node
}

func newNode(v interface{}) *node {
	return &node{v, nil, nil}
}

type List struct {
	size int
	head *node
	tail *node
}

func newList() *List {
	list := &List{0, newNode(0), newNode(0)}
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

func (l *List) Empty() bool {
	return l.head.next == l.tail
}

func (l *List) PutAtHead(v interface{}) {
	putNode := newNode(v)
	putNode.next = l.head.next
	l.head.next.prev = putNode
	l.head.next = putNode
	putNode.prev = l.head
	l.size++
}

func (l *List) PutAtTail(v interface{}) {
	putNode := newNode(v)
	beforeTail := l.tail.prev
	beforeTail.next = putNode
	putNode.prev = beforeTail
	putNode.next = l.tail
	l.tail.prev = putNode
	l.size++
}

func (l *List) Erase(pos *node) {
	// 链表的话，多写几个变量。不要一直使用好几重指针
	be := pos.prev
	ed := pos.next
	be.next = ed
	ed.prev = be
	l.size--
}

func (l *List) Size() int {
	return l.size
}
