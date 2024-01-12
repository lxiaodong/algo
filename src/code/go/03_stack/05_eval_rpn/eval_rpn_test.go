package eval_rpn

import (
	"fmt"
	"testing"
)

func TestEvalRPN1(t *testing.T) {
	// fmt.Println(EvalRPN1([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(EvalRPN1([]string{"10", "6", "9", "3", "+", "-11", " * ", "/", " * ", "17", "+", "5", "+"}))
}

func TestEvalRPN2(t *testing.T) {
	// fmt.Println(EvalRPN2([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(EvalRPN2([]string{"10", "6", "9", "3", "+", "-11", " * ", "/", " * ", "17", "+", "5", "+"}))
}
