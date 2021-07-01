package parser

import (
	"github.com/kakudo415/modelang/scanner"
)

type Node interface {
	Pos() int
}

type EOF struct {
	Token scanner.Token
}

type Expr interface {
	Node
}

type Operand struct {
	X scanner.Token
}

type BinaryExpr struct {
	Op scanner.Token
	L  Expr
	R  Expr
}

func (e EOF) Pos() int        { return e.Token.Pos }
func (o Operand) Pos() int    { return o.X.Pos }
func (e BinaryExpr) Pos() int { return e.Op.Pos }
