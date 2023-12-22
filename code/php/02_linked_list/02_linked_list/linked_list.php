<?php

class ListNode
{
    public int $val;

    public ListNode|NULL $next;

    public function __construct(int $val = 0, ListNode $next = NULL)
    {
        $this->val = $val;
        $this->next = $next;
    }
}

/**
 * 设计链表 https://leetcode.cn/problems/design-linked-list/
 */
class LinkedList
{
    public ListNode $dummyHead;
    public int $size;

    public function __construct()
    {
        $this->dummyHead = new ListNode();
        $this->size = 0;
    }


    /**
     * 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
     *
     * @param integer $val
     * @return void
     */
    public function addAtHead(int $val)
    {
        $node = new ListNode($val);
        $node->next = $this->dummyHead->next;
        $this->dummyHead->next = $node;
        $this->size++;
    }

    /**
     * 将值为 val 的节点追加到链表的最后一个元素
     *
     * @param integer $val
     * @return void
     */
    public function addAtTail(int $val)
    {
        $current = $this->dummyHead;
        $node = new ListNode($val);
        while ($current->next != null) {
            $current = $current->next;
        }
        $current->next = $node;
        $this->size++;
    }

    /**
     * 在链表中的第 index 个节点之前添加值为 val  的节点。
     * 如果 index 等于链表的长度，则该节点将附加到链表的末尾。
     * 如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
     *
     * @param integer $index
     * @param integer $val
     * @return void
     */
    public function addAtIndex(int $index, int $val)
    {
        if ($index > $this->size) {
            return;
        }

        if ($index < 0) {
            $index = 0;
        }

        $current = $this->dummyHead;
        while ($index--) {
            $current = $current->next;
        }
        $node = new ListNode($val);
        $node->next = $current->next;
        $current->next = $node;
        $this->size++;
    }

    /**
     * 获取链表中第 index 个节点的值。如果索引无效，则返回-1
     *
     * @param integer $index
     * @return integer
     */
    public function get(int $index): int
    {
        if ($index < 0 || $index >= $this->size) {
            return -1;
        }

        $current = $this->dummyHead->next;
        while ($index--) {
            $current = $current->next;
        }

        return $current->val;
    }

    /**
     * 如果索引 index 有效，则删除链表中的第 index 个节点
     *
     * @param integer $index
     * @return void
     */
    public function deleteAtIndex(int $index)
    {
        if ($index < 0 || $index >= $this->size) {
            return;
        }

        $slow = $this->dummyHead;
        $fast = $this->dummyHead->next;
        while ($index--) {
            $fast = $fast->next;
            $slow = $slow->next;
        }

        $slow->next = $fast->next;
        $this->size--;
    }

    /**
     * 打印
     *
     * @return void
     */
    public function print()
    {
        $current = $this->dummyHead->next;
        while ($current != NULL) {
            print_r($current->val . "\n");
            $current = $current->next;
        }
    }
}

$linkedList = new LinkedList();
$linkedList->addAtHead(1);
$linkedList->addAtTail(3);
$linkedList->addAtIndex(1,2);
$linkedList->print();
var_dump($linkedList->get(1));
$linkedList->deleteAtIndex(1);
$linkedList->print();
var_dump($linkedList->get(1));

