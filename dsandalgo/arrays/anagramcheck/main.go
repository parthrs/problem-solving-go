package main

import (
	"fmt"
	"os"
)

func main() {
	if result := AnagramCheck(os.Args[1], os.Args[2]); result {
		fmt.Printf("%s is an anagram of %s!\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("%s is not an anagram of %s\n", os.Args[1], os.Args[2])
	}
}
