package mathcalculator

import "errors"

// Possible error
var (
	ErrCannotTokenize    = errors.New("Cannot tokenize")
	ErrInvalidOperator   = errors.New("Invalid operator")
	ErrInvalidExpression = errors.New("Invalid expression")
	ErrCannotParseNumber = errors.New("Cannot parse number")
)
