import unittest
from typing import List


class GenerateMatrix:
    """
    螺旋矩阵II https://leetcode.cn/problems/spiral-matrix-ii/
    """

    def generate(self, n: int) -> List[List]:
        # 开始初始值
        startX, startY = 0, 0

        # 中间位置
        mid = n // 2
        
        # 圈数
        loop = n // 2

        # 初始填充数字
        count = 1

        # 初始化数据
        res = [[0] * n for _ in range(n)]

        # 每一层填充元素个数
        offset = 1

        while loop > 0:
            # 上边
            for i in range(startY, n - offset):
                res[startX][i] = count
                count += 1

            # 右边
            for i in range(startX, n - offset):
                res[i][n - offset] = count
                count += 1

            # 下边
            for i in range(n - offset, startY, -1):
                res[n - offset][i] = count
                count += 1

            # 右边
            for i in range(n - offset, startX, -1):
                res[i][startY] = count
                count += 1

            startX += 1
            startY += 1
            offset += 1
            loop -= 1

        if n % 2 != 0:
            res[mid][mid] = count

        return res


class TestGenerateMatrix(unittest.TestCase):
    def test_generate_matrix(self):
        main = GenerateMatrix()
        nums = [[1, 2, 3], [8, 9, 4], [7, 6, 5]]
        self.assertListEqual(main.generate(3), nums)


if __name__ == "__main__":
    unittest.main()
