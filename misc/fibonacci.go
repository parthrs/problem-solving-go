package misc

import "fmt"

func PrintFibonacci() func() {
	i, j := 0, 1
	return func() {
		fmt.Printf("%d\n", i)
		i, j = j, i+j
	}
}
