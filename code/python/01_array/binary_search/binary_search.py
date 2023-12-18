import unittest
from typing import List


class BinarySearch:
    """
    二分查找 https://leetcode.cn/problems/binary-search/
    """

    def search1(self, nums: List[int], target: int) -> int:
        """左闭右闭"""
        if len(nums) == 0:
            return -1

        left = 0
        right = len(nums) - 1

        while left <= right:
            middle = left + (right - left) // 2
            if nums[middle] > target:
                right = middle - 1
            elif nums[middle] < target:
                left = middle + 1
            else:
                return middle

        return -1

    def search2(self, nums: List[int], target: int) -> int:
        """左闭右开"""
        if len(nums) == 0:
            return -1

        left = 0
        right = len(nums)

        while left < right:
            middle = left + (right - left) // 2
            if nums[middle] > target:
                right = middle
            elif nums[middle] < target:
                left = middle + 1
            else:
                return middle

        return -1


class TestBinarySearch(unittest.TestCase):
    def test_search1(self):
        binary_search = BinarySearch()
        nums = [-1, 0, 3, 5, 9, 12]
        target = 9
        self.assertEqual(binary_search.search1(nums, target), 4)

    def test_search2(self):
        binary_search = BinarySearch()
        nums = [-1, 0, 3, 5, 9, 12]
        target = 9
        self.assertEqual(binary_search.search2(nums, target), 4)


if __name__ == "__main__":
    unittest.main()
