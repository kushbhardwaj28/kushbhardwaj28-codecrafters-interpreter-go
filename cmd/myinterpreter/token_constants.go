package main;

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
    TokenType_BAND_EQUAL
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
    TokenType_EOF: "EOF",
    TokenType_LEFT_PAREN: "LEFT_PAREN",
    TokenType_RIGHT_PAREN: "RIGHT_PAREN",
    TokenType_LEFT_BRACE: "LEFT_BRACE",
    TokenType_RIGHT_BRACE: "RIGHT_BRACE",
    TokenType_COMMA: "COMMA",
    TokenType_DOT: "DOT",
    TokenType_MINUS: "MINUS",
    TokenType_PLUS: "PLUS",
    TokenType_SEMICOLON: "SEMICOLOR",
    TokenType_SLASH: "SLASH",
    TokenType_STAR: "STAR",

    TokenType_BANG: "BANG",
    TokenType_BAND_EQUAL: "BAND_EQUAL",
    TokenType_EQUAL: "EQUAL",
    TokenType_EQUAL_EQUAL: "EQUAL_EQUAL",
    TokenType_GREATER: "GREATER",
    TokenType_GREATER_EQUAL: "GREATER_EQUAL",
    TokenType_LESS: "LESS",
    TokenType_LESS_EQUAL: "LESS_EQUAL",

    TokenType_IDENTIFIER: "IDENTIFIER",
    TokenType_STRING: "STRING",
    TokenType_NUMBER: "NUMBER",

    TokenType_AND: "AND",
    TokenType_CLASS: "CLASS",
    TokenType_ELSE: "ELSE",
    TokenType_FALSE: "FLASE",
    TokenType_FUN: "FUN",
    TokenType_FOR: "FOR",
    TokenType_IF: "IF",
    TokenType_NIL: "NIL",
    TokenType_OR: "OR",
    TokenType_PRINT: "PRINT",
    TokenType_RETURN: "RETURN",
    TokenType_SUPER: "SUPER",
    TokenType_THIS: "THIS",
    TokenType_TRUE: "TRUE",
    TokenType_VAR: "VAR",
    TokenType_WHILE: "WHILE",
}
