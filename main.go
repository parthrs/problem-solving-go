package main

import (
	"fmt"

	"github.com/parthrs/problem-solving-go/misc"
	"github.com/parthrs/problem-solving-go/pkg"
)

/*
Go solutions to problems and puzzles.
*/

func main() {
	fmt.Printf("problem-solving-go 🚀\n\n\n")

	s := pkg.NewSetDS[[2]int]()
	fmt.Println(s.Contains([2]int{1, 2}))
	s.Add([2]int{1, 2})
	fmt.Println(s.Contains([2]int{1, 2}))
	fmt.Println(s.Contains([2]int{23, 44}))
	s.Add([2]int{23, 44})
	fmt.Println(s.Contains([2]int{23, 44}))
	fmt.Printf("=========================\n\n")
	misc.PrintHotFuzz()
	fmt.Printf("=========================\n\n")
	misc.ParseLog()
	fmt.Printf("=========================\n\n")
	misc.PrintEmployeeHeirarchy("A123456789", 0)
	fmt.Printf("=========================\n\n")
	p := misc.NewParser()
	_ = p.ParseInput("misc/parsing-rpc-like-messages.txt")
	fmt.Println(p.GetSize("Vehicle"))
	fmt.Println(p.GetSize("Vector64"))
	fmt.Println(p.GetSize("float"))
}
