package main

import "testing"

// Can store all the tests in one data structure
//testMap = [][]interface{
//	[]interface{"", "", true}
//}

var trueAnagrams = [][]string{
	[]string{"old west action", "clint eastwood"},
	[]string{"silent", "listen"},
	[]string{"McDonald's restaurants", "Uncle Sam's standard rot"},
}

var falseAnagrams = [][]string{
	[]string{"tom marvolo riddle", "lord voldemort"},
}

func TestAnagramCheck(t *testing.T) {
	for _, li := range trueAnagrams {
		if result := AnagramCheck(li[0], li[1]); result != true {
			t.Errorf("Expected: true, but got %v\n", result)
		}
	}

	for _, li := range falseAnagrams {
		if result := AnagramCheck(li[0], li[1]); result != false {
			t.Errorf("Expected: false, but got %v\n", result)
		}
	}
}
