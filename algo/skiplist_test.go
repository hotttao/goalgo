package algo

import (
	"testing"
)

type S string

func (s S) Key() string {
	return string(s)
}

func (s S) String() string {
	return string(s)
}

func TestSkipList(t *testing.T) {
	// s := NewSkipList(3, 3)
	// a := S("abc")
	// s.Insert(6, a)
	// // s.String()
	// b := S("find")
	// s.Insert(4, b)
	// // s.String()
	// s.Insert(1, a)
	// s.String()
	// v, ok := s.Find(4)
	// fmt.Printf("%v - %t\n", v.Key(), ok)
	// s.String()
	// s.Delete(4)
	// s.String()
	// s.Insert(10, a)
	// s.Insert(8, a)
	// s.Insert(9, a)
	// s.String()
}
