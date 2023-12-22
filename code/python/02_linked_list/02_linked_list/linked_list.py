import unittest


class ListNode:
    """ "list node"""

    def __init__(self, val: int = 0, next=None) -> None:
        self.val = val
        self.next = next


class LinkedList:
    """
    设计链表 https://leetcode.cn/problems/design-linked-list/
    """

    def __init__(self) -> None:
        self.head = ListNode(0)
        self.size = 0

    def get(self, index: int) -> int:
        """获取链表中第 index 个节点的值。如果索引无效，则返回-1"""
        if index < 0 or index >= self.size:
            return -1

        current = self.head.next
        if current == None:
            return -1

        while index > 0:
            current = current.next
            index -= 1
        return current.val

    def add_at_head(self, val: int):
        """在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。"""
        node = ListNode(val=val)
        node.next = self.head.next
        self.head.next = node
        self.size += 1

    def add_at_tail(self, val: int):
        """将值为 val 的节点追加到链表的最后一个元素"""
        node = ListNode(val=val)
        current = self.head
        while current.next != None:
            current = current.next
        current.next = node
        self.size += 1

    def add_at_index(self, index: int, val: int):
        """
        在链表中的第 index 个节点之前添加值为 val  的节点。
        如果 index 等于链表的长度，则该节点将附加到链表的末尾。
        如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点
        """
        if index < 0 or index > self.size:
            return

        current = self.head
        while index > 0:
            current = current.next
            index -= 1

        node = ListNode(val=val, next=current.next)
        current.next = node
        self.size += 1

    def delete_at_index(self, index: int):
        """如果索引 index 有效，则删除链表中的第 index 个节点"""
        if index < 0 or index >= self.size:
            return

        slow = self.head
        fast = self.head.next

        while index > 0:
            slow = slow.next
            fast = fast.next
            index -= 1

        slow.next = fast.next
        self.size -= 1

    def print(self):
        current = self.head
        while current.next != None:
            print(current.next.val)
            current = current.next


class TestLinkedList(unittest.TestCase):
    def test_linked_list(self):
        linked_list = LinkedList()
        linked_list.add_at_head(1)
        linked_list.add_at_tail(3)
        linked_list.add_at_index(1, 2)
        linked_list.add_at_index(2, 4)
        linked_list.delete_at_index(3)
        linked_list.print()
        print(linked_list.get(1))


if __name__ == "__main__":
    unittest.main()
