package sort

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l-1, r+1
	// must have brace
	tempNum := nums[(l+r)>>1]
	for i < j {
		for ok := true; ok; {
			i++
			if nums[i] >= tempNum {
				ok = false
			}
		}
		for ok := true; ok; {
			j--
			if nums[j] <= tempNum {
				ok = false
			}
		}
		if i < j {
			tempSwap := nums[i]
			nums[i] = nums[j]
			nums[j] = tempSwap
		}
	}
	quickSort(nums, l, j)
	quickSort(nums, j+1, r)
}
