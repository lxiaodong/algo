<?php

/**
 * 有序数组的平方 https://leetcode.cn/problems/squares-of-a-sorted-array/
 */
class SortSquare
{
    /**
     * 有序数组的平方
     *
     * @param array $nums
     * @return array
     */
    public function sortSquare(array $nums): array
    {
        if (empty($nums)) {
            return $nums;
        }

        $res = [];
        for ($i = 0; $i < count($nums); $i++) {
            $res[$i] = 0;
        }

        $index = count($nums) - 1;

        for ($i = 0, $j = count($nums) - 1; $i <= $j;) {
            if ($nums[$i] ** 2 > $nums[$j] ** 2) {
                $res[$index] = $nums[$i] ** 2;
                $i++;
            } else {
                $res[$index] = $nums[$j] ** 2;
                $j--;
            }

            $index--;
        }

        return $res;
    }
}


$sortSquare = new SortSquare();
$nums = [-4, -1, 0, 3, 10];
print_r($sortSquare->sortSquare($nums));

$nums =  [-7, -3, 2, 3, 11];
print_r($sortSquare->sortSquare($nums));