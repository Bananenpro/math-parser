package main

import (
	"fmt"
	"strconv"
)

type TokenType string

const (
	NUMBER      TokenType = "number"
	PLUS        TokenType = "+"
	MINUS       TokenType = "-"
	STAR        TokenType = "*"
	SLASH       TokenType = "/"
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN TokenType = ")"
)

type Token struct {
	Type  TokenType
	Value float64
}

func (t Token) String() string {
	if t.Type == NUMBER {
		return fmt.Sprintf("%v", t.Value)
	}
	return string(t.Type)
}

type ScanError struct {
	Position int
	Message  string
}

func (s ScanError) Error() string {
	return fmt.Sprintf("[%d] %s", s.Position, s.Message)
}

type scanner struct {
	src    []rune
	tokens []Token

	expStartRune int
	currentRune  int
}

func Scan(input string) ([]Token, error) {
	scanner := scanner{
		src:    []rune(input),
		tokens: make([]Token, 1),
	}
	err := scanner.scan()
	return scanner.tokens, err
}

func (e *scanner) scan() error {
	for e.currentRune = 0; e.currentRune < len(e.src); e.currentRune++ {
		c := e.src[e.currentRune]
		switch c {
		case '+':
			e.addToken(PLUS, 0)
		case '-':
			e.addToken(MINUS, 0)
		case '*':
			e.addToken(STAR, 0)
		case '/':
			e.addToken(SLASH, 0)
		case '(':
			e.addToken(LEFT_PAREN, 0)
		case ')':
			e.addToken(RIGHT_PAREN, 0)
		case ' ', '\t', '\r', '\n':
			break
		default:
			if !isNumber(c) {
				return ScanError{
					Position: e.currentRune,
					Message:  fmt.Sprint("invalid character: ", c),
				}
			}
			e.scanNumber()
		}
	}
	return nil
}

func (e *scanner) scanNumber() {
	for e.currentRune < len(e.src) && isNumber(e.src[e.currentRune]) {
		e.currentRune++
	}

	value, _ := strconv.ParseFloat(string(e.src[e.expStartRune:e.currentRune]), 64)

	e.addToken(NUMBER, value)
}

func (e *scanner) addToken(tokenType TokenType, value float64) {
	e.tokens = append(e.tokens, Token{
		Type:  tokenType,
		Value: value,
	})
	e.expStartRune = e.currentRune + 1
}

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}
