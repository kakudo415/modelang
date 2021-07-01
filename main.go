package main

import (
	"bufio"
	"os"

	"github.com/kakudo415/modelang/executor"
	"github.com/kakudo415/modelang/parser"
	"github.com/kakudo415/modelang/scanner"
)

func main() {
	var src []byte
	readline(&src)

	s := new(scanner.Scanner)
	s.Init(src)
	p := new(parser.Parser)
	p.Init(*s)
	e := new(executor.Executor)
	e.Init(*p)
	for {
		err := e.Execute()
		if err != nil {
			break
		}
	}
}

func readline(dest *[]byte) {
	var err error
	reader := bufio.NewReaderSize(os.Stdin, 4096)
	*dest, _, err = reader.ReadLine()
	if err != nil {
		panic("CANNOT READ FROM TERMINAL")
	}
}
