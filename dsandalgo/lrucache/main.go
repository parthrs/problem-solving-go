package main

import "fmt"

func main() {
	c := NewLRUCache(3)
	c.put(1, 100)
	c.put(2, 200)
	c.put(3, 300)
	fmt.Println(c)
	fmt.Println(c.get(1))
	fmt.Println(c)
}
