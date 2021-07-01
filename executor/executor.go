package executor

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/kakudo415/modelang/parser"
	"github.com/kakudo415/modelang/scanner"
)

type Executor struct {
	parser  parser.Parser
	current parser.Node
}

func (e *Executor) Init(p parser.Parser) {
	e.parser = p
	e.forward()
}

func (e *Executor) Execute() error {
	switch n := e.current.(type) {
	case parser.EOF:
		fmt.Printf("EOF\n")
		return errors.New("EOF")
	case parser.Expr:
		fmt.Printf("%f\n", e.calcExpr(n))
	default:
		fmt.Printf("%d\n", n.Pos())
		panic("UNKNOWN NODE")
	}
	e.forward()
	return nil
}

func (e *Executor) forward() {
	e.current = e.parser.Parse()
}

func (e *Executor) calcExpr(root parser.Expr) float64 {
	switch expr := root.(type) {
	case parser.BinaryExpr:
		return e.calcBinaryExpr(expr)
	case parser.Operand:
		num, err := strconv.ParseFloat(expr.X.Raw, 64)
		if err != nil {
			panic(err)
		}
		return num
	default:
		panic("UNKNOWN EXPRESSION")
	}
}

func (e *Executor) calcBinaryExpr(root parser.BinaryExpr) float64 {
	switch root.Op.Type {
	case scanner.ADD:
		ans := e.calcExpr(root.L) + e.calcExpr(root.R)
		return ans
	case scanner.SUB:
		ans := e.calcExpr(root.L) - e.calcExpr(root.R)
		return ans
	case scanner.MUL:
		ans := e.calcExpr(root.L) * e.calcExpr(root.R)
		return ans
	case scanner.QUO:
		ans := e.calcExpr(root.L) / e.calcExpr(root.R)
		return ans
	default:
		panic("UNKNOWN BINARY OPERATOR")
	}
}
