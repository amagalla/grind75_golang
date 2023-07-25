package easy

/*
	Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
	determine if the input string is valid.

	An input string is valid if:

	Open brackets must be closed by the same type of brackets.
	Open brackets must be closed in the correct order.
	Every close bracket has a corresponding open bracket of the same type.


	Example 1:

	Input: s = "()"
	Output: true
	Example 2:

	Input: s = "()[]{}"
	Output: true
	Example 3:

	Input: s = "(]"
	Output: false
*/

func ValidParentheses (s string) bool {
	var stack []string

	cache := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, character := range s {
		char := string(character)

		if val, ok := cache[char]; ok {
			if len(stack) > 0 && val == stack[len(stack) - 1] {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) <= 0
}