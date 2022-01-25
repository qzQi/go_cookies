package stack

import (
	"testing"
)

func TestStackPut(t *testing.T) {
	// st := Stack{0, nil}
	st := Stack{0, nil}
	for i := 0; i < 10; i++ {
		st.Put(make([]int, 10))
	}
	// t.Log("begin test ")
	if st.S() != 10 {
		t.Error("size should be 10")
		t.Fatal("fatal err")
	}
}

func TestGet(t *testing.T) {
	st := Stack{0, nil}

	for i := 0; i < 10; i++ {
		st.Put(i)
	}

	for i := 9; i >= 0; i-- {
		if v := st.Pop(); v != i {
			t.Errorf("get value err when %d\n", i)
		}
	}

	if st.Size != 0 {
		t.Error("size err")
	}
}
