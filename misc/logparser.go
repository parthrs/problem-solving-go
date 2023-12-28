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

/*
This implementation uses a map to keep track of
the count of messages seen in the log file for every
minute entry seen

The key in this map is the min timestamp in string form
and the value is a int counter
*/

func ParseLog() {
	fd, err := os.Open("misc/logparser.log")
	if err != nil {
		fmt.Printf("Error opening file %v. Exiting..", err)
		return
	}
	defer fd.Close()

	reader := bufio.NewScanner(fd)
	reader.Split(bufio.ScanLines)

	// Dec  3 00:03, i.e. up to min accuracy
	logMap := map[string]int{}

	for reader.Scan() {
		text := reader.Text()
		dateTime := strings.Fields(text)                                            // Split the entire line with spaces
		minute := strings.Split(dateTime[2], ":")                                   // Parse the timestamp
		ts := strings.Join(dateTime[:2], " ") + " " + strings.Join(minute[:2], ":") // Combine month day and timestamp (min accuracy)
		logMap[ts]++                                                                // e.g. Dec 3 00:02 or Dec 3 23:59
	}
	fmt.Println("minute,number_of_messages")
	for key, val := range logMap {
		fmt.Printf("%v, %d\n", key, val)
	}
}
