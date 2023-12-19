import unittest
from typing import List


class ListNode:
    def __init__(self, val: int = 0, next=None) -> None:
        self.val = val
        self.next = next


class LinkedList:
    """
    移除链表元素 https://leetcode.cn/problems/remove-linked-list-elements/
    """

    def new_list(self, nums: List[int]) -> ListNode:
        """创建链表"""
        dummy_head = ListNode()
        current = dummy_head
        for num in nums:
            node = ListNode(val=num)
            current.next = node
            current = current.next

        return dummy_head.next
    
    def remove_element(self, list:ListNode, val:int) ->ListNode:
        """移除元素"""
        
        dummy_head = ListNode(next=list)
        current = dummy_head
        while current != None and current.next != None:
            if current.next.val == val:
                current.next = current.next.next
            else:
                current = current.next
        
        return dummy_head.next
    

class TestRemoveElement(unittest.TestCase):
    def test_remove_element(self):
        Linked_list = LinkedList()
        l = Linked_list.new_list([1,2,6,3,4,5,6])
        rml = Linked_list.remove_element(l, 6)
        while rml != None:
            print(rml.val)
            rml = rml.next

if __name__ == "__main__":
    unittest.main()