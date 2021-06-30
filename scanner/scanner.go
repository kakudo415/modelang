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
)

func (s *Scanner) Init(src []byte) {
	s.src = src
	s.forward()
}

func (s *Scanner) Scan() []Token {
	var ts []Token
	for !s.isEOF() {
		switch {
		case s.isSpace():
			s.forward()
		case s.isNumber():
			ts = append(ts, s.scanNumber())
		case s.ch == '+':
			ts = append(ts, Token{Type: ADD, Pos: s.pos, Raw: "+"})
			s.forward()
		case s.ch == '-':
			ts = append(ts, Token{Type: SUB, Pos: s.pos, Raw: "-"})
			s.forward()
		case s.ch == '*':
			ts = append(ts, Token{Type: MUL, Pos: s.pos, Raw: "*"})
			s.forward()
		case s.ch == '/':
			ts = append(ts, Token{Type: QUO, Pos: s.pos, Raw: "/"})
			s.forward()
		default:
			panic("UNKNOWN TOKEN")
		}
	}
	ts = append(ts, Token{Type: EOF, Pos: s.pos})
	return ts
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
