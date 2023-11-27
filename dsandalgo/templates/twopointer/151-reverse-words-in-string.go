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
