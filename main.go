package main

import (
	"github.com/kakudo415/modelang/parser"
	"github.com/kakudo415/modelang/scanner"
)

func main() {
	s := new(scanner.Scanner)
	s.Init([]byte("4 + 3 * (2 - 1) / 5 + 2 - 3 / 6 + 5"))
	p := new(parser.Parser)
	p.Init(*s)
	expr := p.Parse()
	showExpr(expr)
}

func showExpr(root parser.Node) {
	print(root.Token.Raw)
	for _, child := range root.Child {
		showExpr(child)
	}
}
