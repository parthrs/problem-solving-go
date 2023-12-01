package misc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Parse a log and extract the number of time per minute a program writes to the log

Write a script which parses /var/log/messages and generates a CSV with two columns:minute,number_of_messages in sorted time order.

minute,number_of_messages
Dec  3 00:02,1
Dec  3 00:03,6
*/

func ParseLog() {
	fd, err := os.Open("misc/logparser.log") // Get a file descriptor
	if err != nil {
		fmt.Printf("Error opening file %v. Exiting..", err)
		return
	}
	defer fd.Close()

	reader := bufio.NewScanner(fd) // Create a reader around the fd
	reader.Split(bufio.ScanLines)  // Set the reader to split on new lines

	// Dec  3 00:03:05
	logMap := map[string]int{}

	for reader.Scan() { // Scan returns true till EOL
		text := reader.Text()            // Text fetches text, if scan gave true
		dateTime := strings.Fields(text) // Fields is split but for whitespaces
		minute := strings.Split(dateTime[2], ":")
		ts := strings.Join(dateTime[:2], " ") + " " + strings.Join(minute[:2], ":") // Join returns strings and Go supports a "+"
		logMap[ts]++
	}
	fmt.Println("minute,number_of_messages")
	for key, val := range logMap {
		fmt.Printf("%v, %d\n", key, val)
	}
}
