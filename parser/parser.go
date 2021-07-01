package parser

import "github.com/kakudo415/modelang/scanner"

type Parser struct {
	scanner scanner.Scanner
	current scanner.Token
}

func (p *Parser) Init(s scanner.Scanner) {
	p.scanner = s
	p.forward()
}

func (p *Parser) Parse() Node {
	switch p.current.Type {
	case scanner.EOF:
		return EOF{Token: p.current}
	case scanner.NUMBER, scanner.LPAREN, scanner.ADD, scanner.SUB:
		return p.parseExpr()
	default:
		panic("UNKNOWN TOKEN")
	}
}

func (p *Parser) forward() {
	p.current = p.scanner.Scan()
}

func (p *Parser) parseExpr() Expr {
	return p.parseBinaryExpr_ADD_SUB()
}

func (p *Parser) parseBinaryExpr_ADD_SUB() Expr {
	l := p.parseBinaryExpr_MUL_QUO()
	for p.current.Type == scanner.ADD || p.current.Type == scanner.SUB {
		var e BinaryExpr
		e.Op = p.current
		p.forward()
		r := p.parseBinaryExpr_MUL_QUO()
		e.L = l
		e.R = r
		l = e
	}
	return l
}

func (p *Parser) parseBinaryExpr_MUL_QUO() Expr {
	l := p.parseOperand()
	for p.current.Type == scanner.MUL || p.current.Type == scanner.QUO {
		var e BinaryExpr
		e.Op = p.current
		p.forward()
		r := p.parseOperand()
		e.L = l
		e.R = r
		l = e
	}
	return l
}

func (p *Parser) parseOperand() Expr {
	if p.current.Type == scanner.LPAREN {
		p.forward()
		e := p.parseExpr()
		if p.current.Type != scanner.RPAREN {
			panic("PAREN IS NOT MATCHED")
		}
		p.forward()
		return e
	}

	var unary string
	if p.current.Type == scanner.ADD || p.current.Type == scanner.SUB {
		unary = p.current.Raw
		p.forward()
	}
	token := p.current
	token.Raw = unary + token.Raw
	p.forward()
	return Operand{X: token}
}
