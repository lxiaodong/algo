package binary_search

// 二分查找 leecode https://leetcode.cn/problems/binary-search/

// 左闭右闭
// 时间复杂度:O(n)
func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	if target < nums[left] {
		return -1
	}

	right := len(nums) - 1
	if target > nums[right] {
		return -1
	}

	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}

	return -1
}

// 左闭右开
// 时间复杂度:O(n)
func BinarySearch1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	if target < nums[left] {
		return -1
	}

	right := len(nums)

	for left < right {
		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}

	return -1
}
