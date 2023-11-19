package pkg

import (
	"reflect"
	"testing"
)

func TestSetDS(t *testing.T) {
	s1 := NewSetDS[[2]int]()
	s2 := NewSetDS[[2]int]()
	s1.Add([2]int{1, 2})
	s1.Add([2]int{3, 4})
	s1.Add([2]int{1, 99})
	s1.Add([2]int{2, 2})
	s1.Add([2]int{4, 4})

	s2.Add([2]int{4, 4})
	s2.Add([2]int{1, 2})
	s2.Add([2]int{100, 200})
	s2.Add([2]int{300, 299})
	s2.Add([2]int{100, 2})

	got := Intersection(s1, s2)
	expected := NewSetDS[[2]int]()
	expected.Add([2]int{1, 2})
	expected.Add([2]int{4, 4})

	if equal := reflect.DeepEqual(got, expected); !equal {
		t.Errorf("Expected: %v, Got:%v\n", expected, got)
	}

	got = Union(s1, s2)
	expected = NewSetDS[[2]int]()
	expected.Add([2]int{1, 2})
	expected.Add([2]int{3, 4})
	expected.Add([2]int{1, 99})
	expected.Add([2]int{2, 2})
	expected.Add([2]int{4, 4})
	expected.Add([2]int{100, 200})
	expected.Add([2]int{300, 299})
	expected.Add([2]int{100, 2})

	if equal := reflect.DeepEqual(got, expected); !equal {
		t.Errorf("Expected: %v, Got:%v\n", expected, got)
	}
}
