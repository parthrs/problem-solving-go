package main

import (
	"fmt"
	"time"

	"github.com/parthrs/LetsGo/basics"
	"github.com/parthrs/LetsGo/dsandalgo/pkg"
	"github.com/parthrs/LetsGo/misc"
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
	fmt.Printf("=========================\n\n")
	s := pkg.NewSetDS[[2]int]()
	fmt.Println(s.Contains([2]int{1, 2}))
	s.Add([2]int{1, 2})
	fmt.Println(s.Contains([2]int{1, 2}))
	fmt.Println(s.Contains([2]int{23, 44}))
	s.Add([2]int{23, 44})
	fmt.Println(s.Contains([2]int{23, 44}))
	fmt.Printf("=========================\n\n")
	k := map[string]int{
		"one":         1,
		"one-hundred": 100,
		"two":         2,
		"ninety-nine": 99,
	}
	basics.PrintMapOnSortedValues(k)
	fmt.Printf("=========================\n\n")
	misc.PrintHotFuzz()
	fmt.Printf("=========================\n\n")
	misc.ParseLog()
	fmt.Printf("=========================\n\n")
	misc.PrintEmployeeHeirarchy("A123456789", 0)
	fmt.Printf("=========================\n\n")
	// The "outer func" SomeSetup runs first, then the "inner func" it returns runs after time.Sleep
	defer basics.SomeSetup()()
	time.Sleep(3 * time.Second)
	fmt.Printf("=========================\n\n")
}
