<?php

/**
 * 螺旋矩阵II https://leetcode.cn/problems/spiral-matrix-ii/
 */
class GenerateMatrix
{

    /**
     * 创建螺旋矩阵
     * 例如：输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
     *
     * @param integer $n
     * @return array
     */
    public function generate(int $n): array
    {
        // 填充数组
        $res = array_fill(0, $n, array_fill(0, $n, 0));
        
        // 中间位置
        $mid = floor($n / 2);

        // 每个圈循环几次
        $loop = floor($n / 2);

        // 开始坐标
        $startX = $startY = 0;

        // 每一层填充元素个数
        $offset = 1;

        // 每个位置填充的数字
        $count = 1;

        while ($loop > 0) {
            $i = $startX;
            $j = $startY;
            
            // 上层
            for (; $j < $n - $offset; $j++) {
                $res[$i][$j] = $count++;
            }

            // 右边
            for (; $i < $n - $offset; $i++) {
                $res[$i][$j] = $count++;
            }

            // 下边
            for (; $j > $startY; $j--) {
                $res[$i][$j] = $count++;
            }

            // 左边
            for (; $i > $startX; $i--) {
                $res[$i][$j] = $count++;
            }

            $startX++;
            $startY++;
            $offset++;
            $loop--;
        }

        if ($n % 2 == 1) {
            $res[$mid][$mid] = $count;
        }

        return $res;
    }
}

$main = new GenerateMatrix();
var_dump($main->generate(3));