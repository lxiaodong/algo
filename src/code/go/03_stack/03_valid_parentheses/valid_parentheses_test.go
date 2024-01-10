package valid_parentheses

import (
	"fmt"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	fmt.Println(IsValid1("([)]"))
	fmt.Println(IsValid2("()"))
}
