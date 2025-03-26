package main

import "fmt"

type IToken interface {
	String() string
}

func (tt TokenType) String() string {
	if str, exists := TokenTypeName[tt]; exists {
		return str
	}
	return "unknown"
}

type Token struct {
	_type   TokenType
	lexeme  string
	literal any
	line    int
}

func NewToken(_type TokenType, lexeme string, literal any, line int) *Token {
	return &Token{
		_type:   _type,
		lexeme:  lexeme,
		literal: literal,
		line:    line,
	}
}

func (t *Token) String() string {
	var literal = "null"
	if t.literal != nil {
		switch v := t.literal.(type) {
		case string:
			literal = v
		case []byte:
			literal = string(v)
		default:
			literal = fmt.Sprintf("%v", v)
		}
	}
	var typeStr string = t._type.String()

	return fmt.Sprintf("%s %s %s", typeStr, t.lexeme, literal)
}
