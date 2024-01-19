<?php

require_once "./stack_interface.php";

/**
 * 数组栈
 */
class ArrayStack implements StackInterface
{
    /** array 存放数据 */
    private array $data;
    /** int 栈顶 */
    private int $top;

    public function __construct()
    {
        $this->top = -1;
    }

    /**
     * 添加元素
     *
     * @param mixed $val
     * @return void
     */
    public function push(mixed $val)
    {
        $this->top++;
        $this->data[$this->top] = $val;
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

        $val = $this->data[$this->top];
        $this->top--;
        return $val;
    }

    /**
     * 判断是否为空
     *
     * @return boolean
     */
    public function isEmpty(): bool
    {
        return $this->top() == -1;
    }

    /**
     * 清空栈
     *
     * @return void
     */
    public function flush()
    {
        $this->top = -1;
    }

    /**
     * 返回栈顶元素
     *
     * @return mixed
     */
    public function top(): mixed
    {
        return $this->data[$this->top];
    }
}

$stack = new ArrayStack();
$stack->push(1);
$stack->push("a");
print($stack->top());
print($stack->pop());
print($stack->pop());
