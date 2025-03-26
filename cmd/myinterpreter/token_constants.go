package main

type TokenType int

const (
	TokenType_EOF TokenType = iota
	TokenType_LEFT_PAREN
	TokenType_RIGHT_PAREN
	TokenType_LEFT_BRACE
	TokenType_RIGHT_BRACE
	TokenType_COMMA
	TokenType_DOT
	TokenType_MINUS
	TokenType_PLUS
	TokenType_SEMICOLON
	TokenType_SLASH
	TokenType_STAR

	TokenType_BANG
	TokenType_BANG_EQUAL
	TokenType_EQUAL
	TokenType_EQUAL_EQUAL
	TokenType_GREATER
	TokenType_GREATER_EQUAL
	TokenType_LESS
	TokenType_LESS_EQUAL

	TokenType_IDENTIFIER
	TokenType_STRING
	TokenType_NUMBER

	TokenType_AND
	TokenType_CLASS
	TokenType_ELSE
	TokenType_FALSE
	TokenType_FUN
	TokenType_FOR
	TokenType_IF
	TokenType_NIL
	TokenType_OR
	TokenType_PRINT
	TokenType_RETURN
	TokenType_SUPER
	TokenType_THIS
	TokenType_TRUE
	TokenType_VAR
	TokenType_WHILE
)

var TokenTypeName = map[TokenType]string{
	TokenType_EOF:         "EOF",
	TokenType_LEFT_PAREN:  "LEFT_PAREN",
	TokenType_RIGHT_PAREN: "RIGHT_PAREN",
	TokenType_LEFT_BRACE:  "LEFT_BRACE",
	TokenType_RIGHT_BRACE: "RIGHT_BRACE",
	TokenType_COMMA:       "COMMA",
	TokenType_DOT:         "DOT",
	TokenType_MINUS:       "MINUS",
	TokenType_PLUS:        "PLUS",
	TokenType_SEMICOLON:   "SEMICOLON",
	TokenType_SLASH:       "SLASH",
	TokenType_STAR:        "STAR",

	TokenType_BANG:          "BANG",
	TokenType_BANG_EQUAL:    "BANG_EQUAL",
	TokenType_EQUAL:         "EQUAL",
	TokenType_EQUAL_EQUAL:   "EQUAL_EQUAL",
	TokenType_GREATER:       "GREATER",
	TokenType_GREATER_EQUAL: "GREATER_EQUAL",
	TokenType_LESS:          "LESS",
	TokenType_LESS_EQUAL:    "LESS_EQUAL",

	TokenType_IDENTIFIER: "IDENTIFIER",
	TokenType_STRING:     "STRING",
	TokenType_NUMBER:     "NUMBER",

	TokenType_AND:    "AND",
	TokenType_CLASS:  "CLASS",
	TokenType_ELSE:   "ELSE",
	TokenType_FALSE:  "FALSE",
	TokenType_FUN:    "FUN",
	TokenType_FOR:    "FOR",
	TokenType_IF:     "IF",
	TokenType_NIL:    "NIL",
	TokenType_OR:     "OR",
	TokenType_PRINT:  "PRINT",
	TokenType_RETURN: "RETURN",
	TokenType_SUPER:  "SUPER",
	TokenType_THIS:   "THIS",
	TokenType_TRUE:   "TRUE",
	TokenType_VAR:    "VAR",
	TokenType_WHILE:  "WHILE",
}

var SymbolTokenType = map[string]TokenType{
	"":  TokenType_EOF,
	"(": TokenType_LEFT_PAREN,
	")": TokenType_RIGHT_PAREN,
	"{": TokenType_LEFT_BRACE,
	"}": TokenType_RIGHT_BRACE,
	",": TokenType_COMMA,
	".": TokenType_DOT,
	"-": TokenType_MINUS,
	"+": TokenType_PLUS,
	";": TokenType_SEMICOLON,
	"/": TokenType_SLASH,
	"*": TokenType_STAR,

	"!":  TokenType_BANG,
	"!=": TokenType_BANG_EQUAL,
	"=":  TokenType_EQUAL,
	"==": TokenType_EQUAL_EQUAL,
	">":  TokenType_GREATER,
	">=": TokenType_GREATER_EQUAL,
	"<":  TokenType_LESS,
	"<=": TokenType_LESS_EQUAL,

	//"": TokenType_IDENTIFIER,
	//"": TokenType_STRING,
	//"": TokenType_NUMBER,

	"&&":     TokenType_AND,
	"class":  TokenType_CLASS,
	"else":   TokenType_ELSE,
	"false":  TokenType_FALSE,
	"fun":    TokenType_FUN,
	"for":    TokenType_FOR,
	"if":     TokenType_IF,
	"nil":    TokenType_NIL,
	"or":     TokenType_OR,
	"print":  TokenType_PRINT,
	"return": TokenType_RETURN,
	"super":  TokenType_SUPER,
	"this":   TokenType_THIS,
	"true":   TokenType_TRUE,
	"var":    TokenType_VAR,
	"while":  TokenType_WHILE,
}

var TokenTypeSymbol = map[TokenType]string{
	TokenType_EOF:         "",
	TokenType_LEFT_PAREN:  "(",
	TokenType_RIGHT_PAREN: ")",
	TokenType_LEFT_BRACE:  "{",
	TokenType_RIGHT_BRACE: "}",
	TokenType_COMMA:       ",",
	TokenType_DOT:         ".",
	TokenType_MINUS:       "-",
	TokenType_PLUS:        "+",
	TokenType_SEMICOLON:   ";",
	TokenType_SLASH:       "/",
	TokenType_STAR:        "*",

	TokenType_BANG:          "!",
	TokenType_BANG_EQUAL:    "!=",
	TokenType_EQUAL:         "=",
	TokenType_EQUAL_EQUAL:   "==",
	TokenType_GREATER:       ">",
	TokenType_GREATER_EQUAL: ">=",
	TokenType_LESS:          "<",
	TokenType_LESS_EQUAL:    "<=",

	//TokenType_IDENTIFIER: "",
	//TokenType_STRING: "",
	//TokenType_NUMBER: "",

	TokenType_AND:    "&&",
	TokenType_CLASS:  "class",
	TokenType_ELSE:   "else",
	TokenType_FALSE:  "false",
	TokenType_FUN:    "fun",
	TokenType_FOR:    "for",
	TokenType_IF:     "if",
	TokenType_NIL:    "nil",
	TokenType_OR:     "or",
	TokenType_PRINT:  "print",
	TokenType_RETURN: "return",
	TokenType_SUPER:  "super",
	TokenType_THIS:   "this",
	TokenType_TRUE:   "true",
	TokenType_VAR:    "var",
	TokenType_WHILE:  "while",
}

var MultiCharLexemes = []TokenType{TokenType_BANG, TokenType_LESS, TokenType_GREATER, TokenType_EQUAL}
