package lib001

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a, b := 10, 11
	ans := Add(a, b)
	if ans != 21 {
		t.Fatal("dead")
	}
}
