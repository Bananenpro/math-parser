package main

import (
	"fmt"
	"os"
)

type parser struct {
	current int

	tokens []Token
}

func Parse(tokens []Token) (Expr, error) {
	parser := parser{
		current: 0,
		tokens:  tokens,
	}

	return parser.expression(), nil
}

func (p *parser) expression() Expr {
	return p.term()
}

func (p *parser) term() Expr {
	expr := p.factor()

	for p.match(PLUS, MINUS) {
		operator := p.previous()
		right := p.factor()
		expr = Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}

func (p *parser) factor() Expr {
	expr := p.unary()

	for p.match(STAR, SLASH) {
		operator := p.previous()
		right := p.unary()
		expr = Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}

func (p *parser) unary() Expr {
	if p.match(MINUS) {
		operator := p.previous()
		right := p.unary()
		return Unary{
			Operator: operator,
			Right:    right,
		}
	}

	return p.primary()
}

func (p *parser) primary() Expr {
	if p.match(LEFT_PAREN) {
		expr := p.expression()
		if !p.match(RIGHT_PAREN) {
			fmt.Println("Expect ')' after '('.")
			os.Exit(1)
		}
		return Grouping{
			Expr: expr,
		}
	}

	return Literal{
		Value: p.advance().Value,
	}
}

func (p *parser) match(expected ...TokenType) bool {
	for _, t := range expected {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *parser) check(expected TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == expected
}

func (p *parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

func (p *parser) peek() Token {
	return p.tokens[p.current]
}

func (p *parser) previous() Token {
	return p.tokens[p.current-1]
}
