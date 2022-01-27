package list

import (
	"testing"
)

func TestPutAtHead(t *testing.T) {
	testList := newList()
	if testList.Size() != 0 {
		t.Error("init list size should be zero\n")
	}
	for i := 0; i < 10; i++ {
		testList.PutAtHead(i)
	}
	be := testList.head.next
	for i := 9; be != testList.tail && i >= 0; i, be = i-1, be.next {
		if be.val != i {
			t.Errorf("node->val:%v\t should be:%v\n", be.val, i)
		}
	}
	if testList.Size() != 10 {
		t.Error("after put size should be 10")
	}
}
