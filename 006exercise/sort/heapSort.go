package sort

func down(nums []int, pos, ed int) {
	i := pos
	if 2*pos+1 <= ed && nums[2*pos+1] < nums[i] {
		i = 2*pos + 1
	}
	if 2*pos+2 <= ed && nums[2*pos+2] < nums[i] {
		i = 2*pos + 2
	}
	if i != pos {
		tempSwap := nums[i]
		nums[i] = nums[pos]
		nums[pos] = tempSwap
		down(nums, i, ed)
	}
}

func heapSort(nums []int, l, r int) {
	ed := len(nums) - 1
	if ed <= 0 {
		return
	}

	// make heap
	for i := ed / 2; i >= 0; i-- {
		down(nums, i, ed)
	}

	for i := 0; i < len(nums); i++ {
		tempSwap := nums[0]
		nums[0] = nums[ed]
		nums[ed] = tempSwap
		ed--
		down(nums, 0, ed)
	}
}
