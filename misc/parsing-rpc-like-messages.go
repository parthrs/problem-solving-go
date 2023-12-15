package misc

import (
	"bufio"
	"os"
	"strings"
)

/*
Say you have RPC like messages that you wish to parse from a file
and convert into a queryable object.

string -> 256 bytes, int -> 4 bytes and float -> 4 bytes

For a file containing the following

Message Vehicle
string -> "Hello"
float -> 5.6
float -> 4
int -> 0

Message Vector64
string -> "Right"
float -> 2.2

get_size("Vehicle") -> 256 + 4 + 4 + 4 -> 272
get_size("float") -> 4
get_size("Vector64") -> 8
*/

type Parser struct {
	Messages map[string]int
}

func NewParser() *Parser {
	return &Parser{
		Messages: map[string]int{
			"float":  4,
			"int":    4,
			"string": 256,
		},
	}
}

func (p *Parser) ParseInput(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	messageType := ""

	for scanner.Scan() {
		line := scanner.Text()
		lineFields := strings.Fields(line)

		if len(lineFields) == 0 {
			messageType = ""
			continue
		}

		switch lineFields[0] {
		case "Message":
			messageType = lineFields[1]
		case "float":
			p.Messages[messageType] += 4
		case "int":
			p.Messages[messageType] += 4
		case "string":
			p.Messages[messageType] += 264
		}
	}
	return nil
}

func (p *Parser) GetSize(messageType string) int {
	if val, found := p.Messages[messageType]; found {
		return val
	}
	return -1
}
