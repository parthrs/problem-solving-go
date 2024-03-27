package main

import (
	"context"
	"fmt"
	"net/http"
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

func PackageInstaller() {
	serv := misc.NewPackageServer()
	err := serv.Start()
	time.Sleep(time.Millisecond * 300)
	if err != nil {
		fmt.Printf("Unable to start package server: %v\n", err)
		fmt.Printf("Skipping..")
	} else {
		//time.Sleep(60 * time.Second)
		req, err := http.NewRequest("GET", "http://127.0.0.1:37899/addnode/serverA", nil)
		if err != nil {
			fmt.Printf("Unable to create request to add node on package server: %v\n", err)
			return
		}
		client := http.Client{
			Timeout: time.Second * 5,
		}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error making request to package server: %v\n", err)
			return
		}
		if resp.StatusCode != 200 {
			fmt.Printf("Non-ok response from package server: %v\n", resp.Status)
			return
		}

		g := misc.NewPackage("G", []*misc.Package{})
		f := misc.NewPackage("F", []*misc.Package{})
		i := misc.NewPackage("I", []*misc.Package{})
		d := misc.NewPackage("D", []*misc.Package{})

		c := misc.NewPackage("C", []*misc.Package{g, f, i, d})
		e := misc.NewPackage("E", []*misc.Package{c})

		h := misc.NewPackage("H", []*misc.Package{e, d})
		b := misc.NewPackage("B", []*misc.Package{h})

		misc.InstallPackage(b, "serverA")
		fmt.Println("Sleeping, please verify install...") // curl http://127.0.0.1:37899/getpackages/serverA
		time.Sleep(time.Second * 120)
		fmt.Println("Done, shutting down server..")
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

func PrintDeckOfCardMethods() {
	d := misc.NewDeck()
	fmt.Println("========OG deck========")
	d.Print()
	fmt.Println("========OG deck========")
	c := d.GetCards(5)
	fmt.Println("========get 5 cards========")
	for _, card := range c {
		fmt.Println(card.Card)
	}
	fmt.Println("========get 5 cards========")
	fmt.Println("========after get cards deck========")
	d.Print()
	fmt.Println("========after get cards deck========")
	d.PutBackCard(c[0])
	fmt.Println("========after put back Ace of Clud card========")
	d.Print()
	fmt.Println("========after put back Ace of Clud card========")
	d.Shuffle()
	fmt.Println("========after shuffle deck========")
	d.Print()
	fmt.Println("========after shuffle deck========")
	d.Sort()
	fmt.Println("========after sort deck========")
	d.Print()
	fmt.Println("========after sort deck========")
}

func main() {
	fmt.Printf("problem-solving-go 🚀\n\n\n")

	Sets()
	fmt.Printf("=========================\n\n")
	misc.PrintHotFuzz()
	fmt.Printf("=========================\n\n")
	misc.ParseLog()
	fmt.Printf("=========================\n\n")
	EmployeeTraversal()
	fmt.Printf("=========================\n\n")
	RpcLikeParser()
	fmt.Printf("=========================\n\n")
	SinglyLinkedListTruthTable()
	fmt.Printf("=========================\n\n")
	Fibonacci(7)
	fmt.Printf("=========================\n\n")
	fmt.Printf("Fastest bipedal dinosaurs (https://www.mydinosaurs.com/blog/top-10-fastest-dinosaurs-ever-lived-earth/)\n")
	misc.PrintBipedalDinos(false)
	fmt.Printf("=========================\n\n")
	fortuneTeller := misc.NewFortunes("misc/fortune.txt")
	fortuneTeller.TellMeAFortune()
	fortuneTeller.TellMeAFortune()
	fortuneTeller.TellMeAFortune()
	fortuneTeller.TellMeAFortune()
	fmt.Printf("=========================\n\n")
	PackageInstaller()
	fmt.Printf("=========================\n\n")
	PrintDeckOfCardMethods()
}
