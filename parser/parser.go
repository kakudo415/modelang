package parser

import "github.com/kakudo415/modelang/scanner"

type Parser struct {
	scanner scanner.Scanner
	current scanner.Token
}

type Node struct {
	Token scanner.Token
	Child []Node
}

func (p *Parser) Init(s scanner.Scanner) {
	p.scanner = s
	p.forward()
}

func (p *Parser) Parse() Node {
	switch p.current.Type {
	case scanner.EOF:
		return Node{Token: p.current}
	case scanner.NUMBER:
		return p.parseExprLv2()
	default:
		panic("UNKNOWN TOKEN")
	}
}

func (p *Parser) forward() {
	p.current = p.scanner.Scan()
}

func (p *Parser) parseExprLv2() (n Node) {
	left := p.parseExprLv1()
	for p.current.Type == scanner.ADD || p.current.Type == scanner.SUB {
		n = Node{Token: p.current}
		p.forward()
		right := p.parseExprLv1()
		n.Child = append(n.Child, left)
		n.Child = append(n.Child, right)
		left = n
	}
	return left
}

func (p *Parser) parseExprLv1() (n Node) {
	left := p.parseExprLv0()
	for p.current.Type == scanner.MUL || p.current.Type == scanner.QUO {
		n = Node{Token: p.current}
		p.forward()
		right := p.parseExprLv0()
		n.Child = append(n.Child, left)
		n.Child = append(n.Child, right)
		left = n
	}
	return left
}

func (p *Parser) parseExprLv0() (n Node) {
	if p.current.Type != scanner.LPAREN {
		defer p.forward()
		return Node{Token: p.current}
	}
	p.forward()
	n = p.parseExprLv2()
	if p.current.Type != scanner.RPAREN {
		panic("PAREN IS NOT MATCHED")
	}
	p.forward()
	return n
}
