package stacks

import (
	"strings"
)

/*
Given a string s which represents an expression, evaluate this expression and return its value.

The integer division should truncate toward zero.

1 <= s.length <= 3 * 10^5
s consists of integers and operators ('+', '-', '*', '/') separated by some number of spaces.
s represents a valid expression.
All the integers in the expression are non-negative integers in the range [0, 2^31 - 1].
The answer is guaranteed to fit in a 32-bit integer.
*/

// Push addition and subtractions to stack; i.e. <sign><num> -> stack e.g. -2
// Solve multiplication/div operations with top of stack, push result to stack
// Finally loop over stack to +- the results into final result
// Since we only push to stack (the presign and number) when we encounter
// an operand, the 'index == len(s) - 1' ensures the last num gets handled.
func Calculate(s string) (res int) { // This logic has us doing running total in the end, so initialize res here helps
	num := 0
	preSign := '+' // Assuming first num in expression never negative
	stack := []int{}

	s = strings.Trim(s, " ") // Trim leading and trailing whitespaces
	for index, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0') // num *10 is to account for multiple digits of num, int(ch-'0') uses rune val for 0 as base
		}
		if !isDigit && ch != ' ' || index == len(s)-1 { // if we use else here, we'd have to bend over backwards to tackle the last character
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num // Pop, multiply, push result back
			default:
				stack[len(stack)-1] /= num // Pop, divide, push result back
			}
			preSign = ch // Update presign
			num = 0      // Reset num pointer
		}
	}

	for _, num := range stack {
		res += num
	}

	return
}
