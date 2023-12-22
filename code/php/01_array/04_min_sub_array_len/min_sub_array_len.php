<?php

use Random\Engine\Mt19937;

/**
 * 最小子数组长度 https://leetcode.cn/problems/minimum-size-subarray-sum/
 */
class MinSubArrayLen
{
    /**
     * 获取最小子数组长度
     *
     * @param array $nums
     * @param integer $val
     * @return integer
     */
    public function minSubArrayLen(array $nums, int $val): int
    {
        if (count($nums) == 0) {
            return 0;
        }
        $fast = $slow = $length = $sum = 0;
        
        while ($fast < count($nums)) {
            $sum += $nums[$fast];

            while($sum >= $val) {
                $length = $fast - $slow + 1;
                $sum -= $nums[$slow];
                $slow++;
            }

            $fast++;
        }

        return $length;
    }
}

$main = new MinSubArrayLen();
$nums =  [2,3,1,2,4,3];
$val = 7;
print($main->minSubArrayLen($nums, $val));