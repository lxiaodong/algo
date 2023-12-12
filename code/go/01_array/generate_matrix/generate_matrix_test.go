package generate_matrix

import (
	"fmt"
	"testing"
)

// 测试螺旋矩阵
func TestGenerateMatrix(t *testing.T) {
	n := 4
	res := GenerateMatrix(n)
	fmt.Println(res)
}
