import unittest


class ListNode:
    def __init__(self, val: int = 0, next=None) -> None:
        self.val = val
        self.next = next


class LinkedList:
    """
    删除链表的倒数第N个节点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
    """

    def __init__(self) -> None:
        self.head = ListNode()
        self.size = 0

    def add_at_tail(self, val: int):
        current = self.head

        while current.next != None:
            current = current.next

        current.next = ListNode(val=val)
        self.size += 1

    def remove_nth_from_end1(self, n: int):
        """通过size计算偏移量"""
        current = self.head

        fast = current.next
        slow = current

        idx = self.size - n  # 5 - 1 // 下标4
        while idx > 0:
            fast = fast.next
            slow = slow.next
            idx -= 1

        slow.next = fast.next
        self.size -= 1

        return self.head.next

    def remove_nth_from_end2(self, n: int):
        """双指针方法"""
        fast = self.head
        slow = self.head

        while n > 0 and fast != None:
            fast = fast.next
            n -= 1

        fast = fast.next
        while fast != None:
            fast = fast.next
            slow = slow.next

        slow.next = slow.next.next

    def print(self):
        current = self.head
        while current.next != None:
            print(current.next.val)
            current = current.next


class TestRemoveNthFromEnd(unittest.TestCase):
    def test_remove_nth_from_end1(self):
        list = LinkedList()
        list.add_at_tail(1)
        list.add_at_tail(2)
        # list.add_at_tail(3)
        # list.add_at_tail(4)
        # list.add_at_tail(5)

        list.remove_nth_from_end1(1)
        list.print()

    def test_remove_nth_from_end2(self):
        list = LinkedList()
        list.add_at_tail(1)
        list.add_at_tail(2)
        list.add_at_tail(3)
        list.add_at_tail(4)
        list.add_at_tail(5)

        list.remove_nth_from_end2(1)
        list.print()


if __name__ == "__main__":
    unittest.main()
