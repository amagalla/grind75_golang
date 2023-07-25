package easy

import (
	"grind75/leetcode/easy"
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct{
		s			string
		expected	bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"", true},
		{")", false},
	}

	for _, test := range tests {
		output := easy.ValidParentheses(test.s)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("ValidParentheses (%s) actual: %t - expected %t",
			test.s, output, test.expected)
		}
	}
}