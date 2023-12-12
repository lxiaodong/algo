package _1_array

// 最小子数组长度 https://leetcode.cn/problems/minimum-size-subarray-sum/

// 粗暴获取最小数组长度， 循环套循环，主要记录最终结果
// 时间复杂度：O(n^2)
func MinSubArrayLen1(nums []int, s int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	res := 0
	minLen := 0
	for i := 0; i < l; i++ {
		c := 0
		for j := i; j < l; j++ {
			c += nums[i]
			if c >= s {
				minLen = j - i + 1
				if res == 0 || minLen < res {
					res = minLen
				}
				break
			}
			c += nums[j]
		}
	}

	return res
}

// 双指针方法(滑动窗口)
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func MinSubArrayLen2(nums []int, s int) int {
	l := len(nums)
	minLen := 0
	res := 0
	c := 0
	i := 0
	for j := 0; j < l; j++ {
		c += nums[j]
		for c >= s {
			minLen = j - i + 1
			if res == 0 || minLen < res {
				res = minLen
			}
			c -= nums[i]
			i++
		}
	}

	return res
}
