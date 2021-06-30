package main

import (
	"fmt"

	"github.com/kakudo415/modelang/scanner"
)

func main() {
	s := new(scanner.Scanner)
	s.Init([]byte("1 + 23*456"))
	ts := s.Scan()
	for _, t := range ts {
		fmt.Println(t.Raw)
	}
}
