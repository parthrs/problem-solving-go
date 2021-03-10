package employee

import "fmt"

// employee struct is not exported
// but the methods (constructor like) New and LeavesRemaining are

type employee struct {
	firstName       string
	lastName        string
	totalLeaves     int
	remainingLeaves int
}

// Go convention
// 'Newemployee' if more types are defined here
func New(firstName string, lastName string, totalLeaves int, remainingLeaves int) employee {
	e := employee{firstName, lastName, totalLeaves, remainingLeaves}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Println("%s %s has %d/%d leaves remaining", e.firstName, e.lastName, e.remainingLeaves, e.totalLeaves)
}
