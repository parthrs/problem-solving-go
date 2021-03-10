package main

import "howtostructs/employee"

func main() {
	e := employee.New()
	f := employee.New("Sam", "Gamgee", 100, 99)

	e.LeavesRemaining()
	f.LeavesRemaining()
}
