package homework

import (
	"testing"
)

// 只出现一次数字 测试用例
func TestSingleNumber(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, item := range tests {
		got := SingleNumber(item.input)
		if got != item.expect {
			t.Errorf("输入%v，预期%d，得到%d", item.input, item.expect, got)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	// 回文数测试用例
	tests := []struct {
		input  int
		expect bool
	}{
		{121, true},
		{-121, false},
		{0, true},
		{12321, true},
		{123, false},
	}

	for _, item := range tests {

		got := IsPalindrome(item.input)

		if got != item.expect {
			t.Errorf("s输入: %d ，预期: %v ，实际得到: %v", item.input, item.expect, got)
		}
	}

}

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"([])", true},
		{"(]", false},
		{"{[(])}", false}, //错误嵌套
		{"(", false},      //左括号多余
		{")", false},      //开头右括号
		{"((()))", true},
	}

	for _, item := range tests {
		got := IsValidParentheses(item.input)
		if got != item.expect {
			t.Errorf("输入字符串：%s，预期：%v，实际得到：%v", item.input, item.expect, got)
		}
	}

}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input  []string
		expect string
	}{
		{[]string{"flott", "floii", "flff"}, "fl"},
		{[]string{"mcctt", "mccii", "mcoff"}, "mc"},
	}

	for _, item := range tests {
		got := LongestCommonPrefix(item.input)
		if got != item.expect {
			t.Errorf("输入%v 预期:%s 实际:%s", item.input, item.expect, got)
		}

	}

}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{8, 9, 9}, []int{9, 0, 0}},
		{[]int{8, 8, 9}, []int{8, 9, 0}},
	}

	for _, item := range tests {
		got := PlusOne(item.input)

		match := true

		if len(got) != len(item.expect) {
			match = false
		} else {
			for i := range got {

				if got[i] != item.expect[i] {
					match = false
				}

			}
		}

		if !match {
			t.Errorf("输入%v 预期%v 实际%v", item.input, item.expect, got)
		}
	}

}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, item := range tests {
		got := TwoSum2(item.nums, item.target)
		match := false
		//两种顺序都算正确 [a,b] 或 [b,a]
		if (got[0] == item.expect[0] && got[1] == item.expect[1]) ||
			(got[0] == item.expect[1] && got[1] == item.expect[0]) {
			match = true
		}
		if !match {
			t.Errorf("nums:%v target:%d 预期%v 实际%v",
				item.nums, item.target, item.expect, got)
		}
	}
}
