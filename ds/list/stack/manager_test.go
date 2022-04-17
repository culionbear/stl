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

func TestPush(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	for it := s.Begin(); it.End(); it = it.Next() {
		fmt.Println(it.Value())
	}
}

func TestPop(t *testing.T) {
	s := New[int]()
	fmt.Println(s.Pop())
	s.PushMore(1, 2, 3)
	for s.Size() != 0 {
		fmt.Println(s.Pop())
	}
}

func TestPushMore(t *testing.T) {
	s := New[int]()
	s.PushMore(1, 2, 3, 4, 8, 99, -1)
	for it := s.Begin(); it.End(); it = it.Next() {
		fmt.Println(it.Value())
	}
}

func TestSize(t *testing.T) {
	s := New[int]()
	fmt.Println(s.Size())
	s.PushMore(1, 2, 3, 4)
	fmt.Println(s.Size())
}

func TestFirst(t *testing.T) {
	s := New[int]()
	fmt.Println(s.First())
	s.PushMore(1, 2, 3, 4)
	fmt.Println(s.First())

	sta := New[string]()
	sta.PushMore("111", "fang", "bear")
	fmt.Println(sta.First())
}

func TestTop(t *testing.T) {
	sta := New[string]()
	fmt.Println(sta.Top())
	sta.PushMore("111", "fang", "bear")
	fmt.Println(sta.Top())
}

func TestIsEmpty(t *testing.T) {
	s := New[int]()
	fmt.Println(s.IsEmpty())
	s.Push(1)
	fmt.Println(s.IsEmpty())
}
