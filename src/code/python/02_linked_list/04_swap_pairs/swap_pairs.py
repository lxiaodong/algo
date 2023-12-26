import unittest

class ListNode:
    def __init__(self, val: int = 0, next=None) -> None:
        self.val = val
        self.next = next
        
class LinkedList:
    """
    两两交换链表中的节点 https://leetcode.cn/problems/swap-nodes-in-pairs/
    """
    
    def __init__(self) -> None:
        self.head = ListNode()
        self.size = 0
        
    def add_at_tail(self, val:int):
        current = self.head
        
        while current.next != None:
            current = current.next
        
        current.next = ListNode(val=val)
        self.size += 1

    def swap_pairs(self) ->ListNode:
        current = self.head
        
        while current.next != None and current.next.next != None:
            n = current.next
            nn = current.next.next
            
            n.next = nn.next
            nn.next = n
            current.next = nn
            
            current = current.next.next
        
        return self.head.next

class TestSwapPairs(unittest.TestCase):
    def test_swap_pairs(self):
        list = LinkedList()
        list.add_at_tail(1)
        list.add_at_tail(2)
        list.add_at_tail(3)
        list.add_at_tail(4)
        
        res = list.swap_pairs()
        
        while res != None:
            print(res.val)
            res = res.next
            
if __name__ == "__main__":
    unittest.main()