package mathcalculator

import "github.com/vanhtuan0409/rtokenizer"

var (
	number       = rtokenizer.TokenType("Number")
	plus         = rtokenizer.TokenType("Plus")
	minus        = rtokenizer.TokenType("Minus")
	multiply     = rtokenizer.TokenType("Multiply")
	divide       = rtokenizer.TokenType("Divide")
	openBracket  = rtokenizer.TokenType("OpenBracket")
	closeBracket = rtokenizer.TokenType("CloseBracket")
)
