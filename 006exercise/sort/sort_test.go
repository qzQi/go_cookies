package sort

import (
	"testing"
)

func TestQuick(t *testing.T) {
	arr := []int{8, 6, 5, 9, 7, 4, 3, 2, 1, 0}
	quickSort(arr, 0, len(arr)-1)

	for i := 0; i <= 9; i++ {
		if i != arr[i] {
			t.Errorf("in pos %v should be arr[%v]\n", i, arr[i])
		}
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{8, 6, 5, 9, 7, 4, 3, 2, 1, 0}
	mergeSort(arr, 0, len(arr)-1)

	for i := 0; i <= 9; i++ {
		if i != arr[i] {
			t.Errorf("in pos %v should be arr[%v]\n", i, arr[i])
		}
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{8, 6, 5, 9, 7, 4, 3, 2, 1, 0}
	heapSort(arr, 0, len(arr)-1)

	for i := 0; i <= 9; i++ {
		if i != 9-arr[i] {
			t.Errorf("in pos %v should be %v but ac arr[%v]\n", i, 9-i, arr[i])
		}
	}
}
