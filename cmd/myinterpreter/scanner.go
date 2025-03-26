package main

import "fmt"

type IScanner interface {
	scanTokens() []Token
}

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) scanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, *NewToken(TokenType_EOF, "", nil, s.line))
	return s.tokens
}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	var c string = s.advance()

	if _type, exists := SymbolTokenType[c]; exists {
		var containsMultiLexemes = Contains(MultiCharLexemes, _type)
		if containsMultiLexemes && s.match(TokenType_EQUAL) {
			s.addToken(getMutliLexem(_type, TokenType_EQUAL))
		} else {
			s.addToken(_type)
		}
	} else {
		hasError = true
		_logger._Error(fmt.Sprintf("[line %d] Error: Unexpected character: %s", s.line, c))
	}
}

func (s *Scanner) _scanToken() {
	var c string = s.advance()
	switch c {
	case "(":
		s.addToken(TokenType_LEFT_PAREN)
		break
	case ")":
		s.addToken(TokenType_RIGHT_PAREN)
		break
	case "{":
		s.addToken(TokenType_LEFT_BRACE)
		break
	case "}":
		s.addToken(TokenType_RIGHT_BRACE)
		break
	case ",":
		s.addToken(TokenType_COMMA)
		break
	case ".":
		s.addToken(TokenType_DOT)
		break
	case "-":
		s.addToken(TokenType_MINUS)
		break
	case "+":
		s.addToken(TokenType_PLUS)
		break
	case ";":
		s.addToken(TokenType_SEMICOLON)
		break
	case "*":
		s.addToken(TokenType_STAR)
		break
	default:
		hasError = true
		_logger._Error(fmt.Sprintf("[line %d] Error: Unexpected character: %s", s.line, c))
	}
}

func (s *Scanner) advance() string {
	var substr = string([]rune(s.source)[s.current])
	s.current++

	return substr
}

func (s *Scanner) addToken(_type TokenType) {
	s._addToken(_type, nil)
}

func (s *Scanner) _addToken(_type TokenType, literal any) {
	var text string = s.source[s.start:s.current]
	s.tokens = append(s.tokens, *NewToken(_type, text, literal, s.line))
}

func (s *Scanner) match(expectedType TokenType) bool {
	if s.isAtEnd() {
		return false
	}
	var substr = string([]rune(s.source)[s.current])
	if SymbolTokenType[substr] != expectedType {
		return false
	}

	s.current++
	return true
}

func getMutliLexem(oldType TokenType, addType TokenType) TokenType {
	var newCharStr = TokenTypeSymbol[oldType] + TokenTypeSymbol[addType]

	if _type, exists := SymbolTokenType[newCharStr]; exists {
		return _type
	} else {
		return TokenType_EOF
	}
}
