<?php

require_once "./stack_interface.php";

class LinkedListNode
{
    public function __construct(
        public mixed $val = 0,
        public LinkedListNode|NULL $next = null
    ) {
    }
}

class LinkedListStack implements StackInterface
{
    /** LinkedNode 栈顶结点 */
    public LinkedListNode|NULL $top = null;

    /**
     * 添加元素
     *
     * @param mixed $val
     * @return void
     */
    public function push(mixed $val)
    {
        $node = new LinkedListNode($val, $this->top);
        $this->top = $node;
    }

    /**
     * 弹出元素
     *
     * @return mixed
     */
    public function pop(): mixed
    {
        if ($this->isEmpty()) {
            return null;
        }

        $node = $this->top;
        $this->top = $this->top->next;
        return $node->val;
    }

    /**
     * 判断是否为空
     *
     * @return boolean
     */
    public function isEmpty(): bool
    {
        return $this->top == null;
    }

    /**
     * 清空栈
     *
     * @return void
     */
    public function flush()
    {
        $this->top = null;
    }

    /**
     * 返回栈顶元素
     *
     * @return mixed
     */
    public function top(): mixed
    {
        if ($this->isEmpty()) {
            return null;
        }

        return $this->top->val;
    }
}

$stack = new LinkedListStack();
$stack->push(1);
$stack->push("a");
print($stack->top());
print($stack->pop());
print($stack->pop());
