package twopointer

import "strings"

func Reverse(s string) string {
	ss := strings.Fields(s)
	i := 0
	j := len(ss) - 1
	for i < j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
	return strings.Join(ss, " ")
}

// func Reverse(s string) string {
// words := strings.Fields(s)
// n := len(words) - 1
// for n >= 0 {
//     result += words[n]
//     if n >= 1 {
//       result += " "
//     }
//     n--
// }
// return
// }
