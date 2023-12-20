package main

import (
	"context"
	"fmt"
	"time"

	"github.com/parthrs/problem-solving-go/leet-code/linkedlist"
	"github.com/parthrs/problem-solving-go/misc"
	"github.com/parthrs/problem-solving-go/pkg"
)

/*
Go solutions to problems and puzzles.
*/

func Sets() {
	s := pkg.NewSetDS[[2]int]()
	fmt.Println(s.Contains([2]int{1, 2}))
	s.Add([2]int{1, 2})
	fmt.Println(s.Contains([2]int{1, 2}))
	fmt.Println(s.Contains([2]int{23, 44}))
	s.Add([2]int{23, 44})
	fmt.Println(s.Contains([2]int{23, 44}))
}

func EmployeeTraversal() {
	serv := misc.NewHTTPServer()
	ch := make(chan error, 1)
	serv.Start(ch)
	if err := <-ch; err != nil {
		fmt.Printf("Error starting the server %v, skipping the next steps..\n", err)
	} else {
		fmt.Printf("Server up, making request..\n")
		time.Sleep(time.Millisecond * 300)
		misc.PrintEmployeeHeirarchy("A123456789", 0)
		fmt.Printf("Done, stopping server..\n")
		serv.Stop(context.Background())
		time.Sleep(time.Millisecond * 300)
	}
}

func RpcLikeParser() {
	p := misc.NewParser()
	_ = p.ParseInput("misc/parsing-rpc-like-messages.txt")
	fmt.Println(p.GetSize("Vehicle"))
	fmt.Println(p.GetSize("Vector64"))
	fmt.Println(p.GetSize("float"))
}

func SinglyLinkedListTruthTable() {
	// Create a new Singly list
	ll := pkg.NewSinglyList[int]()
	// Add elems
	// nil <- 5 <- 4 <- 3 <- 2 <- 1
	//                            Head
	for i := 5; i > 0; i-- {
		ll.AddNodeHead(i)
	}
	fmt.Printf("odd len, slow=fast=head\n")
	linkedlist.PrintMiddleAndLastNodes[int](ll.Head, ll.Head) // 5, 3
	n := pkg.NewSinglyNode[int](0)
	n.Next = ll.Head
	// nil <- 5 <- 4 <- 3 <- 2 <- 1   <- 0
	//                            Head   n
	fmt.Printf("odd len, slow=head; fast=-1\n")
	linkedlist.PrintMiddleAndLastNodes[int](ll.Head, n) // nil, 4
	fmt.Printf("odd len, slow=fast=-1\n")
	linkedlist.PrintMiddleAndLastNodes[int](n, n) // nil, 3 (same as even len)
	fmt.Printf("odd len, slow=-1; fast=head\n")
	linkedlist.PrintMiddleAndLastNodes[int](n, ll.Head)
	fmt.Printf("even len, slow=fast=head\n")
	ll.AddNodeHead(0)
	// nil <- 5 <- 4 <- 3 <- 2 <- 1 <- 0
	//                                Head
	linkedlist.PrintMiddleAndLastNodes[int](ll.Head, ll.Head) // nil, 3
	n.Value = -1
	n.Next = ll.Head
	fmt.Printf("even len, slow=head; fast=-1\n")
	linkedlist.PrintMiddleAndLastNodes[int](ll.Head, n)
	fmt.Printf("even len, slow=fast=-1\n")
	linkedlist.PrintMiddleAndLastNodes[int](n, n)
	fmt.Printf("even len, slow=-1; fast=head\n")
	linkedlist.PrintMiddleAndLastNodes[int](n, ll.Head)
}

func Fibonacci(n int) {
	f := misc.PrintFibonacci()
	for n > 0 {
		f()
		n--
	}
}

func main() {
	fmt.Printf("problem-solving-go 🚀\n\n\n")

	// Sets()
	// fmt.Printf("=========================\n\n")
	// misc.PrintHotFuzz()
	// fmt.Printf("=========================\n\n")
	// misc.ParseLog()
	// fmt.Printf("=========================\n\n")
	// EmployeeTraversal()
	// fmt.Printf("=========================\n\n")
	// RpcLikeParser()
	// fmt.Printf("=========================\n\n")
	// SinglyLinkedListTruthTable()
	// fmt.Printf("=========================\n\n")
	Fibonacci(7)
	fmt.Printf("=========================\n\n")
}
