package move_item

import "fmt"

// 移除元素 https://leetcode.cn/problems/remove-element/

// 循环遍历
// 时间复杂度：O(n^2)
// 空间复杂度：O(1) , 原地互换
func MoveItem1(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	len := len(nums)
	for i := 0; i < len; i++ {
		if nums[i] == target {
			for j := i + 1; j < len; j++ {
				nums[j-1] = nums[j]
			}
			len--
			i--
		}
	}

	return len
}

// 双指针法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func MoveItem2(nums []int, target int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}

	slowIdx := 0

	for fastIdx := 0; fastIdx < len; fastIdx++ {
		if nums[fastIdx] != target {
			nums[slowIdx] = nums[fastIdx]
			slowIdx++
		}
	}

	return slowIdx
}

// 相向双指针方法，基于元素顺序可以改变的题目描述改变了元素相对位置，确保了移动最少元素
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func MoveItem3(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	leftIdx := 0
	rightIdx := len(nums) - 1
	for leftIdx <= rightIdx {
		for leftIdx <= rightIdx && nums[leftIdx] != target {
			leftIdx++
		}

		for leftIdx <= rightIdx && nums[rightIdx] == target {
			rightIdx--
		}
		if leftIdx <= rightIdx {
			nums[leftIdx] = nums[rightIdx]
			leftIdx++
			rightIdx--
		}
	}
	fmt.Println(nums)
	return leftIdx
}
