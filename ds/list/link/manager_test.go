package link

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	l := New(1, 2, 3)
	for it := l.Begin(); it.End(); it = it.Next() {
		fmt.Println(it.Value())
	}
}
