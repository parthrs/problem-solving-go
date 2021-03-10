package main

import "fmt"

// The idea is to embed one struct into another

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) authorDetails() string {
	return fmt.Sprintf("%s %s: %s", a.firstName, a.lastName, a.bio)
}

type post struct {
	title   string
	content string
	author  // Embedding one struct into another, here name is not needed
}

func (p post) postDetails() {
	fmt.Printf("Title: %s\n", p.title)
	fmt.Printf("Content: %s\n", p.content)
	fmt.Printf("Author: %s\n", p.authorDetails()) // Short for p.author.authorDetails()
	fmt.Printf("Bio: %s\n", p.bio)
}

// Embedding a slice of structs
type blog struct {
	posts []post // Embedded slices need names
}

func (b blog) website() {
	fmt.Println("Blog")
	for _, p := range b.posts {
		fmt.Println(p.postDetails)
	}
}

func main() {
	a := author{
		"Frodo",
		"Baggins",
		"Shire <3, thats all.",
	}

	p := post{
		"Lord of the Rings",
		"It started with the one Ring..",
		a,
	}

	fmt.Println(a.authorDetails())
	fmt.Println()
	p.postDetails()

	w := blog{
		posts: []post{p},
	}
	w.website()
}
