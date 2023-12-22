import unittest


class ListNode:
    def __init__(self, val: int = 0, next=None) -> None:
        self.val = val
        self.next = next


class LinkedList:
    """
    反转链表 https://leetcode.cn/problems/reverse-linked-list/
    """

    def __init__(self) -> None:
        self.head = ListNode(val=0)
        self.size = 0

    def add_at_tail(self, val: int):
        current = self.head
        while current.next != None:
            current = current.next

        node = ListNode(val=val)
        current.next = node
        self.size += 1

    def reverse1(self) -> ListNode:
        """双指针方式"""
        current = self.head.next
        pre = None
        while current != None:
            next = current.next
            current.next = pre
            pre = current
            current = next

        return pre

    def reverse2(self) -> ListNode:
        """递归方式"""
        return self._help(None, self.head)

    def _help(self, pre: ListNode, head: ListNode) -> ListNode:
        if head == None:
            return pre

        next = head.next
        head.next = pre
        return self._help(pre=head, head=next)


class TestReverseList(unittest.TestCase):
    def test_reverse1(self):
        list = LinkedList()
        list.add_at_tail(1)
        list.add_at_tail(2)
        list.add_at_tail(3)
        list.add_at_tail(4)

        res = list.reverse1()
        while res != None:
            print(res.val)
            res = res.next

    def test_reverse2(self):
        list = LinkedList()
        list.add_at_tail(1)
        list.add_at_tail(2)
        list.add_at_tail(3)
        list.add_at_tail(4)

        res = list.reverse2()
        while res.next != None:
            print(res.val)
            res = res.next


if __name__ == "__main__":
    unittest.main()
