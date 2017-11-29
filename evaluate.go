package mathcalculator

import (
	"strconv"

	"github.com/vanhtuan0409/rtokenizer"
)

func evaluateSingle(operator *rtokenizer.Token, operand1 *rtokenizer.Token, operand2 *rtokenizer.Token) (*rtokenizer.Token, error) {
	num1, err := getNumberValue(operand1)
	num2, err := getNumberValue(operand2)
	if err != nil {
		return nil, err
	}

	var result float64
	if operator.Type == plus {
		result = num1 + num2
	} else if operator.Type == minus {
		result = num1 - num2
	} else if operator.Type == multiply {
		result = num1 * num2
	} else if operator.Type == divide {
		result = num1 / num2
	} else {
		return nil, ErrInvalidOperator
	}

	return &rtokenizer.Token{
		Type:     number,
		RawValue: strconv.FormatFloat(result, 'f', 2, 64),
	}, nil
}

// Algorithm on: https://stackoverflow.com/questions/13421424/how-to-evaluate-an-infix-expression-in-just-one-scan-using-stacks
func evaluate(tokens []*rtokenizer.Token) (float64, error) {
	postfixTokens, err := convertToPostfix(tokens)
	if err != nil {
		return 0, err
	}

	operandStack := tokenStack{}
	for _, token := range postfixTokens {
		// If token is number
		if token.Type == number {
			operandStack.Push(token)
			continue
		}

		// If token is operator
		operand2 := operandStack.Pop()
		operand1 := operandStack.Pop()
		if operand1 == nil || operand2 == nil {
			return 0, ErrInvalidExpression
		}
		result, err := evaluateSingle(token, operand1, operand2)
		if err != nil {
			return 0, err
		}
		operandStack.Push(result)
	}

	if len(operandStack) != 1 {
		return 0, ErrInvalidExpression
	}
	result := operandStack.Pop()
	return getNumberValue(result)
}

// Evaluate Evaluate an math expression
func Evaluate(expression string) (float64, error) {
	t := rtokenizer.NewTokenizer(rtokenizer.Option{
		IgnoreSpaces:    true,
		IgnoreLineBreak: true,
	})
	t.Add(number, `[0-9]*\.?[0-9]+`)
	t.Add(plus, `\+`)
	t.Add(minus, `-`)
	t.Add(multiply, `\*`)
	t.Add(divide, `/`)
	t.Add(openBracket, `\(`)
	t.Add(closeBracket, `\)`)

	tokens, err := t.Tokenize(expression)
	if err != nil {
		return 0, ErrCannotTokenize
	}

	return evaluate(tokens)
}
