package vector

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestIterator(t *testing.T) {
	v := New[int]()
	v.PushMore(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for it := v.Begin(); it != v.End(); it = it.Next() {
		fmt.Println(it.Value())
	}
}

func TestPopBack(t *testing.T) {
	v := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	v.PopBack()
	v.PopBack()
	v.PopBack()
	v.PopBack()
	v.PopBack()
	v.PopBack()
}

func TestAt(t *testing.T) {
	v := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("v.At(1): %v\n", v.At(1))
	fmt.Printf("v.At(11): %v\n", v.At(11))
	fmt.Printf("v.At(-2): %v\n", v.At(-2))
	fmt.Printf("v.At(0): %v\n", v.At(0))
	fmt.Printf("v.At(-10): %v\n", v.At(-10))
}

func TestRemove(t *testing.T) {
	v := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("v.Remove(-10): %v\n", v.Remove(-10))
	fmt.Printf("v.Remove(1): %v\n", v.Remove(1))
	fmt.Printf("v.Remove(-1): %v\n", v.Remove(-1))
	fmt.Printf("v.Size(): %v\n", v.Size())
	fmt.Printf("v.cap: %v\n", v.cap)
	for it := v.Begin(); it != v.End(); it = it.Next() {
		fmt.Println(it.Value())
	}
}

func TestBack(t *testing.T) {
	v := New(1, 2, 3)
	fmt.Printf("v.Back(): %v\n", v.Back())
}

func TestFront(t *testing.T) {
	v := New(1, 2, 3)
	fmt.Printf("v.Front(): %v\n", v.Front())
}

func TestClear(t *testing.T) {
	v := New(1, 2, 3)
	v.Clear()
	fmt.Printf("v.cap: %v\n", v.cap)
	fmt.Printf("v.Size(): %v\n", v.Size())
}

func TestReverse(t *testing.T) {
	v := New(1, 777, 23, 31)
	v.Reverse()
	for it := v.Begin(); it != v.End(); it = it.Next() {
		fmt.Print(it.Value(), " ")
	}
}

func TestSort(t *testing.T) {
	v := New(1, 777, 23, 31)
	v.Sort(func(i1, i2 int) bool { return i1 < i2 })
	for it := v.Begin(); it != v.End(); it = it.Next() {
		fmt.Print(it.Value(), " ")
	}
}

func TestInsert(t *testing.T) {
	v := New(1, 2, 3)
	v.Insert(-3, 999)
	for it := v.Begin(); it != v.End(); it = it.Next() {
		fmt.Print(it.Value(), " ")
	}
}

func BenchmarkPush(b *testing.B) {
	v := New[int]()
	for i := 0; i < b.N; i++ {
		v.Push(1)
	}
}

func BenchmarkPushMore(b *testing.B) {
	v := New[int]()
	for i := 0; i < b.N; i++ {
		v.PushMore(1, 2, 3, 4, 5)
	}
}

func BenchmarkPopBack(b *testing.B) {
	v := New(1, 2, 3)
	for i := 0; i < b.N; i++ {
		v.PopBack()
	}
}

func TestSort1(t *testing.T) {
	v := New(1, 2, 3)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000000; i++ {
		v.Push(rand.Intn(500000000))
	}
	s := time.Now().Unix()
	v.Sort(func(i1, i2 int) bool { return i1 < i2 })
	e := time.Now().Unix()
	fmt.Println(e - s)
}
