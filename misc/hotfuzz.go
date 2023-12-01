package misc

/*
Loop through numbers 1 to 100, for a number that is divisible by 3, print 'Hot'.
For a number that is divisible by 6, print 'Fuzz'. For number that is divisible by both, print 'HotFuzz'.
For other numbers just print the numbers.
*/

import "fmt"

func PrintHotFuzz() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d ", i)
		if (i%3 == 0) && (i%6 == 0) {
			fmt.Printf("HotFuzz")
		} else if i%3 == 0 {
			fmt.Printf("Hot")
		} else if i%6 == 0 {
			fmt.Printf("Fuzz")
		}
		fmt.Printf("\n")
	}
}
