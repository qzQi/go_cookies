package sort

// import "testing"

var tempMerge [10000]int

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	i, j := l, mid+1
	k := 0
	for i <= mid && j <= r {
		if nums[i] < nums[j] {
			tempMerge[k] = nums[i]
			i, k = i+1, k+1
		} else {
			tempMerge[k] = nums[j]
			j, k = j+1, k+1
		}
	}
	for i <= mid {
		tempMerge[k] = nums[i]
		k, i = k+1, i+1
	}
	for j <= r {
		tempMerge[k] = nums[j]
		j, k = j+1, k+1
	}
	for i, k = l, 0; i <= r; i, k = i+1, k+1 {
		nums[i] = tempMerge[k]
	}
}
