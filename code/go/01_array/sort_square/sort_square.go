package sort_square

// 有序数组的平方 https://leetcode.cn/problems/squares-of-a-sorted-array/

// 双指针法
// 时间复杂度：O(1)
func SortSquare1(nums []int) []int {
	len := len(nums)
	l, r, i := 0, len-1, len-1
	res := make([]int, len)
	for l <= r {
		ln, rn := nums[l]*nums[l], nums[r]*nums[r]
		if ln < rn {
			res[i] = rn
			r--
		} else {
			res[i] = ln
			l++
		}
		i--
	}
	return res
}
