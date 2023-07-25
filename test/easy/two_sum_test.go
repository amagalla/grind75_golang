package easy

import (
	"grind75/leetcode/easy"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{1, 2, 3, 4}, 10, []int{-1, -1}},
	}

	for _, test := range tests {
		output := easy.TwoSum(test.nums, test.target)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TwoSum(%v, %d) actual: %v expected: %v",
				test.nums, test.target, output, test.expected)
		}
	}
}
