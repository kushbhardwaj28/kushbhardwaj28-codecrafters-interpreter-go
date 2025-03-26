package main

import "slices"

func Contains[T comparable](arr []T, item T) bool {
	return slices.Contains(arr, item)
}

func getMutliLexem(oldType TokenType, addType TokenType) TokenType {
	var newCharStr = TokenTypeSymbol[oldType] + TokenTypeSymbol[addType]

	if _type, exists := SymbolTokenType[newCharStr]; exists {
		return _type
	} else {
		return TokenType_EOF
	}
}

