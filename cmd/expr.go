package main

type Expr interface {
	Accept(visitor ExprVisitor) (any, error)
}

type ExprVisitor interface {
	VisitBinary(expr Binary) (any, error)
	VisitGrouping(expr Grouping) (any, error)
	VisitLiteral(expr Literal) (any, error)
	VisitUnary(expr Unary) (any, error)
}

type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (b Binary) Accept(visitor ExprVisitor) (any, error) {
	return visitor.VisitBinary(b)
}

type Grouping struct {
	Expr Expr
}

func (g Grouping) Accept(visitor ExprVisitor) (any, error) {
	return visitor.VisitGrouping(g)
}

type Literal struct {
	Value float64
}

func (l Literal) Accept(visitor ExprVisitor) (any, error) {
	return visitor.VisitLiteral(l)
}

type Unary struct {
	Operator Token
	Right    Expr
}

func (u Unary) Accept(visitor ExprVisitor) (any, error) {
	return visitor.VisitUnary(u)
}
