<?php

class ListNode
{
    public function __construct(
        public int $val,
        public ListNode|NULL $next
    ) {
    }
}

/**
 * 反转链表 https://leetcode.cn/problems/reverse-linked-list/
 */
class LinkedList
{
    public ListNode $head;
    public int $size;

    public function __construct()
    {
        $this->head = new ListNode(0);
        $this->size = 0;
    }

    public function addAtTail(int $val)
    {
        $current = $this->head;
        while ($current->next != Null) {
            $current = $current->next;
        }
        $node = new ListNode($val);
        $current->next = $node;
        $this->size++;
    }

    /**
     * 双指针方法
     *
     * @return ListNode
     */
    public function reverse1(): ListNode
    {
        $current = $this->head->next;
        $pre = null;
        while ($current != NUll) {
            $next = $current->next;
            $current->next = $pre;
            $pre = $current;
            $current = $next;
        }

        return $pre;
    }

    /**
     * 递归法
     *
     * @return ListNode
     */
    public function reverse2(): ListNode
    {
        return $this->help(NULL, $this->head->next);
    }

    private function help(ListNode|NULL $pre, ListNode $head): ListNode
    {
        if ($head == NULL) {
            return $pre;
        }
        $next = $head->next;
        $head->next = $pre;
        return $this->help($head, $next);
    }
}
