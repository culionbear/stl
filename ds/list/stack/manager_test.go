package stack

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	s := New(1, 2, 3)
	for it := s.Begin(); it.End(); it = it.Next() {
		fmt.Println(it.Value())
	}
}
