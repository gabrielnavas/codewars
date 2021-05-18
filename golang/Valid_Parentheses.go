package main

import "fmt"

func ValidParentheses(parens string) bool {
	if len(parens) == 0 {
		return true
	}
	parensStack := []rune{rune(parens[0])}
	for i := 1; i < len(parens); i++ {
		if len(parensStack) > 0 && parensStack[len(parensStack)-1] == '(' && parens[i] == ')' {
			parensStack = parensStack[0 : len(parensStack)-1]
		} else {
			parensStack = append(parensStack, rune(parens[i]))
		}
	}

	return len(parensStack) == 0
}

func main() {
	fmt.Println(ValidParentheses("()") == true)
	fmt.Println(ValidParentheses(")") == false)
	fmt.Println(ValidParentheses("()") == true)
	fmt.Println(ValidParentheses(")(()))") == false)
	fmt.Println(ValidParentheses("(") == false)
	fmt.Println(ValidParentheses("(())((()())())") == true)
	fmt.Println(ValidParentheses(")))))))))))))") == false)
}
