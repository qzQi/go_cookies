/*
这里主要是练习使用go的string相关的函数：
go的string类似于char slice但是string不可变
需要把他转化为byte slice/rune
*/
package leetcode

// 这里如果可以使用泛型就好了，或则标准库提供
// func max(i,j *int)int{
// 	if *i>*j{
// 		return *j
// 	}
// 	return *i
// }
// 不是因为copy by value
func lengthOfLongestSubstring(s string) int {
	ans := 0
	arr := [300]int{}
	for i, j := 0, 0; j < len(s); j++ {
		arr[int(s[j])]++
		for arr[int(s[j])] > 1 {
			arr[int(s[i])]--
			i++
		}
		temp := j - i + 1
		if temp > ans {
			ans = temp
		}
	}
	return ans
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	// 没有do while还真不会用go写快排呢
	// for ok:=true;ok;{
	// }模拟do while
	i, j := l-1, r+1
	temp := nums[i+j>>1]
	for i < j {
		for ok := true; ok; {
			i++
			if nums[i] <= temp {
				ok = false
			}
		}
		for ok := true; ok; {
			j--
			if nums[j] >= temp {
				ok = false
			}
		}
		if i < j {
			swap := nums[i]
			nums[i] = nums[j]
			nums[j] = swap
		}
	}
	quickSort(nums, l, j)
	quickSort(nums, j+1, r)
}

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[k-1]
}
