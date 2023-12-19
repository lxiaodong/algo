import unittest
from typing import List


class MinSubArrayLen:
    """
    最小子数组长度 https://leetcode.cn/problems/minimum-size-subarray-sum/
    """

    def min_sub_array_len(self, nums: List[int], val:int) ->int:
        """
        获取最小子数组长度
        通过快慢指针
        """
        
        if len(nums) == 0:
            return 0
        
        fast, slow, length, sum = 0, 0, 0, 0
        while fast < len(nums):
            sum += nums[fast]
            while sum >= val:
                length = fast - slow + 1
                sum -= nums[slow]
                slow += 1
                
            fast += 1
            
        return length
        

class TestMinSubArrayLen(unittest.TestCase):
    """单元测试"""
    
    def test_min_sub_array_len(self):
        min_sub_array_len = MinSubArrayLen()
        s = 7 
        nums = [2,3,1,2,4,3]
        self.assertEqual(min_sub_array_len.min_sub_array_len(nums, s), 2)
        

if __name__ == "__main__":
    unittest.main()