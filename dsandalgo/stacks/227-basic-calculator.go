package stacks

import (
	"fmt"
	"strconv"
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

/*
Note: The funcs with suffix 'ValientAttempt' are an incorrect approach
      that I came up with, kept here as a humility kick
*/

func ValientAttemptSolveExpression(number int, expression string) int {
	operator := string(expression[0])
	rhs, _ := strconv.Atoi(string(expression[1]))
	switch operator {
	case "+":
		return number + rhs
	case "-":
		return number - rhs
	case "*":
		return number * rhs
	case "/":
		return number / rhs
	default:
		return 0
	}
}

func ValientAttemptCalculate(s string) int {
	operationOrderMap := map[string]int{
		"*": 1,
		"/": 1,
		"+": 2,
		"-": 2,
	}
	s = strings.ReplaceAll(s, " ", "")
	executionMap := map[int][]string{}

	//fmt.Println(s)
	for i := 1; i < len(s); i++ {
		prevLetter := string(s[i-1])
		letter := string(s[i])
		if _, found := operationOrderMap[letter]; !found {
			order := operationOrderMap[prevLetter]
			executionMap[order] = append(executionMap[order], fmt.Sprintf("%s%s", prevLetter, letter))
		}
	}

	runningResult, _ := strconv.Atoi(string(s[0]))

	for _, expression := range executionMap[1] {
		runningResult = ValientAttemptSolveExpression(runningResult, expression)
	}

	for _, expression := range executionMap[2] {
		runningResult = ValientAttemptSolveExpression(runningResult, expression)
	}

	return runningResult
}

func Calculate(s string) (res int) { // This logic has us doing running total in the end, so initialize res here helps
	num := 0
	preSign := '+' // Assuming first num in expression never negative
	stack := []int{}

	s = strings.Trim(s, " ") // Trim leading and trailing whitespaces
	for index, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = int(ch - '0')
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
