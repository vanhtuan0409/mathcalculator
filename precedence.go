package mathcalculator

import (
	"github.com/vanhtuan0409/rtokenizer"
)

func getOperatorPriority(operator *rtokenizer.Token) int {
	if operator.Type == multiply || operator.Type == divide {
		return 2
	}
	if operator.Type == plus || operator.Type == minus {
		return 1
	}
	return 0
}

func compareOperator(op1 *rtokenizer.Token, op2 *rtokenizer.Token) int {
	return getOperatorPriority(op1) - getOperatorPriority(op2)
}
