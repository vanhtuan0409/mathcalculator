package mathcalculator

import (
	"errors"

	"github.com/vanhtuan0409/rtokenizer"
)

// Algorithm on: http://www.includehelp.com/c/infix-to-postfix-conversion-using-stack-with-c-program.aspx
func convertToPostfix(tokens []*rtokenizer.Token) ([]*rtokenizer.Token, error) {
	result := []*rtokenizer.Token{}
	operatorStack := tokenStack{}
	for _, token := range tokens {
		// If number
		if token.Type == number {
			result = append(result, token)
			continue
		}

		// If open bracket
		if token.Type == openBracket {
			operatorStack.Push(token)
			continue
		}

		// If close bracket
		if token.Type == closeBracket {
			topOperator := operatorStack.Pop()
			if topOperator == nil {
				return nil, errors.New("Unmatch bracket")
			}
			for topOperator.Type != openBracket {
				result = append(result, topOperator)
				topOperator = operatorStack.Pop()
			}
			continue
		}

		// If operator
		topOperator := operatorStack.Pop()
		for topOperator != nil && compareOperator(topOperator, token) >= 0 {
			result = append(result, topOperator)
			topOperator = operatorStack.Pop()
		}
		operatorStack.Push(topOperator)
		operatorStack.Push(token)
	}

	// Push rest of stack to result
	topOperator := operatorStack.Pop()
	for topOperator != nil {
		result = append(result, topOperator)
		topOperator = operatorStack.Pop()
	}

	return result, nil
}
