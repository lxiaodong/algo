import unittest
from typing import List


class SortSquare:
    """
    有序数组的平方 https://leetcode.cn/problems/squares-of-a-sorted-array/
    """

    def sort_square(self, nums: List[int]) -> List[int]:
        """有序数组的平方"""

        l, r, i = 0, len(nums) - 1, len(nums) - 1
        res = [float("inf")] * len(nums)  # 需要提前定义列表，存放结果

        while l <= r:
            if nums[l] ** 2 < nums[r] ** 2:
                res[i] = nums[r] ** 2
                r -= 1
            else:
                res[i] = nums[l] ** 2
                l += 1

            i -= 1

        return res


class TestSortSquare(unittest.TestCase):
    def test_sort_square(self):
        sort_square = SortSquare()
        nums = [-4, -1, 0, 3, 10]
        self.assertListEqual(sort_square.sort_square(nums), [0, 1, 9, 16, 100])

        nums = [-7, -3, 2, 3, 11]
        self.assertListEqual(sort_square.sort_square(nums), [4, 9, 9, 49, 121])


if __name__ == "__main__":
    unittest.main()
