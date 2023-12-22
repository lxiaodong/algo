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
 * 移除链表元素 https://leetcode.cn/problems/remove-linked-list-elements/
 */
class LinkedList
{
    /**
     * 创建链表
     *
     * @param array $nums
     * @return ListNode
     */
    public function newList(array $nums): ListNode
    {
        $dummyHeader = new ListNode();
        $current = $dummyHeader;
        foreach ($nums as $key => $val) {
            $node = new ListNode($val);
            $current->next = $node;
            $current = $current->next;
        }
        return $dummyHeader->next;
    }

    /**
     * 移除元素
     *
     * @param ListNode $list
     * @param integer $val
     * @return ListNode
     */
    public function removeItem(ListNode $list, int $val): ListNode
    {
        $dummyHeader = new ListNode(next:$list);
        $current = $dummyHeader;
        while ($current != NULL && $current->next != NULL) {
            if ($current->next->val == $val) {
                $current->next = $current->next->next;
            } else {
                $current = $current->next;
            }
        }

        return $dummyHeader->next;
    }
}

$linkedList = new LinkedList();
$list = $linkedList->newList([1, 2, 6, 3, 4, 5, 6]);
$val = 6;
$res = $linkedList->removeItem($list, $val);
while($res != NULL) {
    print_r($res->val."\n");
    $res = $res->next;
}
