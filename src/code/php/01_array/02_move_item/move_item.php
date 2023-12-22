<?php

/**
 * 移除元素 https://leetcode.cn/problems/remove-element/
 */
class MoveItem
{
    /**
     * 移除元素(粗暴移动)
     *
     * @param array $nums
     * @param integer $val
     * @return integer
     */
    public function moveItem1(array $nums, int $val): int
    {
        $length = count($nums);
        for ($i = 0; $i < $length; $i++) {
            if ($nums[$i] == $val) {
                for ($j = $i+1; $j < $length; $j++) {
                    $nums[$j-1] = $nums[$j];
                }
                $length--;
                $i--;
            }
        }

        return $length;
    }

    /**
     * 双指针方法原地移动元素
     *
     * @param array $nums
     * @param integer $val
     * @return integer
     */
    public function moveItem2(array $nums, int $val): int
    {
        $length = count($nums);
        $slow = 0;
        for ($fast = 0; $fast < $length; $fast++) {
            if ($nums[$fast] != $val) {
                $nums[$slow] = $nums[$fast];
                $slow++;
            }
        }
        return $slow;
    }
}

$moveItem = new MoveItem();
$nums = [3,2,2,3]; 
$val = 3;
print($moveItem->moveItem1($nums, $val)."\n");
print($moveItem->moveItem2($nums, $val)."\n");

print("=================\n");

$nums = [0,1,2,2,3,0,4,2]; 
$val = 2;
print($moveItem->moveItem1($nums, $val)."\n");
print($moveItem->moveItem2($nums, $val)."\n");


