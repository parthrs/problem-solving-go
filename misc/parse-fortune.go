package misc

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

/*
Program to parse the fortunes in fortunes.txt,
and then read them out randomly.
*/

type Fortunes struct {
	Unseen []string
	Seen   []string
	File   string
}

func NewFortunes(filename string) *Fortunes {
	f := &Fortunes{
		Unseen: []string{},
		Seen:   []string{},
		File:   filename,
	}
	f.parseFortunes()
	return f
}

func (f *Fortunes) parseFortunes() error {
	file, err := os.Open(f.File)
	if err != nil {
		log.Fatalf("Unable to open fortunes file: %v", err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanRunes)

	fortune := ""
	prev, curr := "", ""

	for scan.Scan() {
		curr = scan.Text()
		if curr == "%" {
			prev = curr
			continue
		} else if curr == "\n" && prev == "%" {
			if fortune != "" {
				f.Unseen = append(f.Unseen, fortune)
			}
			fortune = ""
		} else {
			fortune += curr
			prev = curr
		}
	}
	return nil
}

func (f *Fortunes) TellMeAFortune() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(f.Unseen))
	fortune := f.Unseen[r]
	f.Unseen = append(f.Unseen[:r], f.Unseen[r+1:]...)
	f.Seen = append(f.Seen, fortune)
	fmt.Println(fortune)
}
