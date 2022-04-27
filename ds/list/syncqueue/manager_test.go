package queue

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func PrintQueue[V comparable](q *Manager[V]) {
	for it := q.Begin(); it != q.End(); it = it.Next() {
		fmt.Print(it.Value(), " ")
	}
	fmt.Println()
}

func TestIterator(t *testing.T) {
	q := New(1, 2, 3, 4)
	for it := q.Begin(); it != q.End(); it = it.Next() {
		fmt.Print(it.Value(), " ")
	}
	fmt.Println()
}

func TestPush(t *testing.T) {
	wg.Add(2)
	q := New[int]()
	go func() {
		for i := 0; i < 3000; i++ {
			q.Push(i + 10000)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 3000; i++ {
			q.Push(i + 30000)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("s.Size(): %v\n", q.Size())
}

func TestPushMore(t *testing.T) {
	q := New[int]()
	q.PushMore(1, 2, 8, 9, 11)
	PrintQueue(q)
}

func TestSize(t *testing.T) {
	q := New[int]()
	fmt.Println(q.Size())
	q.PushMore(1, 2, 8, 9, 11)
	fmt.Println(q.Size())
}

func TestFirst(t *testing.T) {
	q := New[int]()
	q.PushMore(1, 2, 8, 9, 11)
	fmt.Println(q.First())
}

func TestPop(t *testing.T) {
	wg.Add(2)
	q := New[int]()
	for i := 0; i < 3000; i++ {
		q.Push(i + 30000)
	}
	go func() {
		for i := 0; i < 1300; i++ {
			fmt.Printf("11111s.Pop(): %v\n", q.Pop())
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1300; i++ {
			fmt.Printf("222222s.Pop(): %v\n", q.Pop())
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("s.Size(): %v\n", q.Size())
}

func TestIsEmpty(t *testing.T) {
	q := New[int]()
	fmt.Println(q.IsEmpty())
	q.PushMore(1, 2, 8)
	fmt.Println(q.IsEmpty())
}

func TestFront(t *testing.T) {
	q := New[int]()
	q.PushMore(1, 2, 8, 9, 11)
	fmt.Println(q.Front())
}

func TestTail(t *testing.T) {
	q := New[int]()
	q.PushMore(1, 2, 8, 9, 11)
	fmt.Println(q.Tail())
}

func BenchmarkPush(b *testing.B) {
	q := New[int]()
	for i := 0; i < b.N; i++ {
		q.Push(1)
	}
}

func BenchmarkPushMore(b *testing.B) {
	q := New[int]()
	for i := 0; i < b.N; i++ {
		q.PushMore(1, 2, 3, 4, 5)
	}
}

func BenchmarkPop(b *testing.B) {
	q := New[int]()
	q.PushMore(1, 2, 8, 9, 11)
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}
