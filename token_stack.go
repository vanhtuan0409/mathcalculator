package mathcalculator

import (
	"github.com/vanhtuan0409/rtokenizer"
)

type tokenStack []*rtokenizer.Token

func (s *tokenStack) Push(t *rtokenizer.Token) {
	*s = append(*s, t)
}

func (s *tokenStack) Top() *rtokenizer.Token {
	if len(*s) == 0 {
		return nil
	}
	stackLength := len(*s)
	return (*s)[stackLength-1]
}

func (s *tokenStack) Pop() *rtokenizer.Token {
	top := s.Top()
	if top != nil {
		*s = (*s)[:len(*s)-1]
	}
	return top
}
