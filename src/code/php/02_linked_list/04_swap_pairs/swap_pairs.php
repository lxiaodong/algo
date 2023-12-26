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
 * 两两交换链表中的节点 https://leetcode.cn/problems/swap-nodes-in-pairs/
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
        while ($current->next != NULL) {
            $current = $current->next;
        }

        $current->next = new ListNode($val);
        $this->size++;
    }

    public function swapPairs(): ListNode
    {
        $current = $this->head;
        while ($current->next != NULL && $current->next->next != NULL) {
            $n = $current->next;
            $nn = $current->next->next;

            $n->next = $nn->next;
            $nn->next = $n;
            $current->next = $nn;

            $current = $current->next->next;
        }

        return $this->head->next;
    }
}

$list = new LinkedList();
$list->addAtTail(1);
$list->addAtTail(2);
$list->addAtTail(3);
$list->addAtTail(4);
$list->addAtTail(5);
$res = $list->swapPairs();

while ($res != NULL) {
    print($res->val . "\n");
    $res = $res->next;
}
