package mathcalculator

import (
	"strconv"

	"github.com/vanhtuan0409/rtokenizer"
)

func getNumberValue(token *rtokenizer.Token) (float64, error) {
	if token.Type != number {
		return 0, ErrCannotParseNumber
	}
	return strconv.ParseFloat(token.RawValue, 64)
}
