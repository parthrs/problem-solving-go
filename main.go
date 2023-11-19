package main

import (
	"fmt"

	"github.com/parthrs/LetsGo/basics"
	"github.com/parthrs/LetsGo/dsandalgo/pkg"
)

/*
Go basics, patterns and solutions
to problems.
*/

func main() {
	fmt.Printf("LetsGo 🚀\n\n\n")
	basics.UnderstandingStrings()
	fmt.Printf("=========================\n\n")
	basics.UnderstandingSorts()
	fmt.Printf("=========================\n\n")
	basics.UnderstandReadingAFile()
	fmt.Printf("=========================\n\n")
	basics.UnderstandMakingHTTPCalls()
	s := pkg.NewSetDS[[2]int]()
	fmt.Println(s.Contains([2]int{1, 2}))
	s.Add([2]int{1, 2})
	fmt.Println(s.Contains([2]int{1, 2}))
	fmt.Println(s.Contains([2]int{23, 44}))
	s.Add([2]int{23, 44})
	fmt.Println(s.Contains([2]int{23, 44}))
}
