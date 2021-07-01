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
	switch e.current.Token.Type {
	case scanner.EOF:
		fmt.Printf("EOF\n")
		return errors.New("EOF")
	case scanner.NUMBER:
		fmt.Printf("%s\n", e.current.Token.Raw)
	case scanner.ADD, scanner.SUB, scanner.MUL, scanner.QUO:
		fmt.Printf("%f\n", e.executeExpr(e.current))
	default:
		println(e.current.Token.Type)
		println(e.current.Token.Pos)
		panic("UNKNOWN NODE")
	}
	e.forward()
	return nil
}

func (e *Executor) forward() {
	e.current = e.parser.Parse()
}

func (e *Executor) executeExpr(root parser.Node) float64 {
	switch root.Token.Type {
	case scanner.NUMBER:
		num, err := strconv.ParseFloat(root.Token.Raw, 64)
		if err != nil {
			panic(err)
		}
		return num
	case scanner.ADD:
		ans := e.executeExpr(root.Child[0]) + e.executeExpr(root.Child[1])
		return ans
	case scanner.SUB:
		ans := e.executeExpr(root.Child[0]) - e.executeExpr(root.Child[1])
		return ans
	case scanner.MUL:
		ans := e.executeExpr(root.Child[0]) * e.executeExpr(root.Child[1])
		return ans
	case scanner.QUO:
		ans := e.executeExpr(root.Child[0]) / e.executeExpr(root.Child[1])
		return ans
	default:
		panic("UNKNOWN NODE")
	}
}
