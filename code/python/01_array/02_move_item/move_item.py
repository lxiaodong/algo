import unittest
from typing import List


class MoveItem:
    """
    移除元素 https://leetcode.cn/problems/remove-element/
    """

    def move_item1(self, nums: List[int], target: int) -> int:
        """粗暴移除元素"""
        i, length = 0, len(nums)

        while i < length:
            if nums[i] == target:
                for j in range(i + 1, length):
                    nums[j - 1] = nums[j]
                length -= 1
                i -= 1
            i += 1

        return length

    def move_item2(self, nums: List[int], target: int) -> int:
        """双指针移除元素"""
        slow, fast, length = 0, 0, len(nums)

        while fast < length:
            if nums[fast] != target:
                nums[slow] = nums[fast]
                slow += 1
            fast += 1

        return slow


class TestMoveItem(unittest.TestCase):
    def test_move_item1(self):
        move_item = MoveItem()
        nums = [3,2,2,3] 
        val = 3
        self.assertEqual(move_item.move_item1(nums, val), 2)
        
        nums = [0,1,2,2,3,0,4,2]
        val = 2
        self.assertEqual(move_item.move_item1(nums, val), 5)

    def test_move_item2(self):
        move_item = MoveItem()
        nums = [3,2,2,3] 
        val = 3
        self.assertEqual(move_item.move_item2(nums, val), 2)
        
        nums = [0,1,2,2,3,0,4,2]
        val = 2
        self.assertEqual(move_item.move_item2(nums, val), 5)


if __name__ == "__main__":
    unittest.main()
