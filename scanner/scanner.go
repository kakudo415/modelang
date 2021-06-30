package scanner

import (
	"unicode"
	"unicode/utf8"
)

type Scanner struct {
	src  []byte
	pos  int
	npos int
	ch   rune
}

type TokenType uint

type Token struct {
	Type TokenType
	Pos  int
	Raw  string
}

const (
	ILLEGAL = iota
	EOF

	NUMBER

	ADD
	SUB
	MUL
	QUO

	LPAREN
	RPAREN
)

func (s *Scanner) Init(src []byte) {
	s.src = src
	s.forward()
}

func (s *Scanner) Scan() Token {
	var t Token
	switch {
	case s.isEOF():
		return Token{Type: EOF, Pos: s.pos}
	case s.isSpace():
		s.forward()
		return s.Scan()
	case s.isNumber():
		t = s.scanNumber()
	case s.ch == '+':
		t = s.scanOp(ADD)
	case s.ch == '-':
		t = s.scanOp(SUB)
	case s.ch == '*':
		t = s.scanOp(MUL)
	case s.ch == '/':
		t = s.scanOp(QUO)
	case s.ch == '(':
		t = s.scanOp(LPAREN)
	case s.ch == ')':
		t = s.scanOp(RPAREN)
	default:
		panic("UNKNOWN TOKEN")
	}
	return t
}

func (s *Scanner) forward() {
	s.pos = s.npos
	if s.isEOF() {
		s.ch = utf8.RuneError
		return
	}
	r, w := utf8.DecodeRune(s.src[s.pos:])
	if r == utf8.RuneError || w == 0 {
		panic("CANNOT SCAN AS UTF-8")
	}
	s.npos += w
	s.ch = r
}

func (s *Scanner) isSpace() bool {
	return s.ch == ' '
}

func (s *Scanner) isNumber() bool {
	return unicode.IsDigit(s.ch)
}

func (s *Scanner) isEOF() bool {
	return s.pos >= len(s.src)
}

func (s *Scanner) scanNumber() Token {
	var t Token
	t.Type = NUMBER
	t.Pos = s.pos
	for s.isNumber() {
		t.Raw += string(s.ch)
		s.forward()
	}
	return t
}

func (s *Scanner) scanOp(tt TokenType) Token {
	var t Token
	t.Type = tt
	t.Pos = s.pos
	t.Raw = string(s.ch)
	s.forward()
	return t
}
