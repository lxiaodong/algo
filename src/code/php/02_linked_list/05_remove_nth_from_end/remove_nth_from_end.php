<?php

class ListNode
{
    public function __construct(
        public int $val = 0,
        public ListNode|NULL $next = NULL
    ) {
    }
}

/**
 * 删除链表的倒数第N个节点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
 */
class LinkedList
{
    public ListNode $head;
    public int $size;
    public function __construct()
    {
        $this->head = new ListNode();
        $this->size = 0;
    }

    public function addAtTail(int $val)
    {
        $current = $this->head;
        while ($current->next != Null) {
            $current = $current->next;
        }

        $current->next = new ListNode($val);
        $this->size++;
    }

    public function removeNthFromEnd1(int $n): ListNode
    {
        $current = $this->head;

        $idx = $this->size - $n;
        while ($idx > 0) {
            $current = $current->next;
            $idx--;
        }
        $current->next  = $current->next->next;

        return $this->head->next;
    }

    public function removeNthFromEnd2(int $n): ListNode
    {
        $fast = $this->head;
        $slow = $this->head;

        while ($n > 0 && $fast->next != NULL) {
            $fast = $fast->next;
            $n--;
        }

        $fast = $fast->next;
        while ($fast != Null) {
            $fast = $fast->next;
            $slow = $slow->next;
        }

        $slow->next = $slow->next->next;

        return $this->head->next;
    }

    public function print()
    {
        $current = $this->head;
        while ($current->next != NULL) {
            print($current->next->val . "\n");
            $current = $current->next;
        }
    }
}

$list = new LinkedList();
$list->addAtTail(1);
$list->addAtTail(2);
$list->addAtTail(3);
$list->addAtTail(4);
$list->addAtTail(5);
$list->removeNthFromEnd2(2);
$list->print();
