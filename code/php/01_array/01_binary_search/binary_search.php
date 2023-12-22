<?php

/**
 * 二分查找 https://leetcode.cn/problems/binary-search/
 */
class BinarySearch
{
    /**
     * 左闭右闭
     *
     * @param array $nums
     * @param integer $target
     * @return integer
     */
    public function search1(array $nums, int $target): int
    {
        if (count($nums) == 0) {
            return -1;
        }

        $left = 0;
        $right = count($nums) - 1;

        while ($left <= $right) {
            $middle = $left + ($right - $left) / 2;
            if ($nums[$middle] > $target) {
                $right = $middle - 1;
            } else if ($nums[$middle] < $target) {
                $left = $middle + 1;
            } else {
                return $middle;
            }
        }

        return -1;
    }

    /**
     * 左闭右开
     *
     * @param array $nums
     * @param integer $target
     * @return integer
     */
    public function search2(array $nums, int $target): int
    {
        if (count($nums) == 0) {
            return -1;
        }

        $left = 0;
        $right = count($nums);

        while ($left < $right) {
            $middle = $left + ($right - $left) / 2;
            if ($nums[$middle] > $target) {
                $right = $middle;
            } else if ($nums[$middle] < $target) {
                $left = $middle + 1;
            } else {
                return $middle;
            }
        }

        return -1;
    }
}

$binarySearch =  new BinarySearch();
$nums = [-1, 0, 3, 5, 9, 12];
$target = 9;
print($binarySearch->search1($nums, $target)."\n");
print($binarySearch->search2($nums, $target));
