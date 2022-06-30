package stack

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestIterator(t *testing.T) {
	s := New(1, 2, 3)
	for it := s.Begin(); it != s.End(); it = it.Next() {
		fmt.Println(it.Value())
	}
}

func TestPush(t *testing.T) {
	wg.Add(2)
	s := New[int]()
	go func() {
		for i := 0; i < 3000; i++ {
			s.Push(i + 10000)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 3000; i++ {
			s.Push(i + 30000)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("s.Size(): %v\n", s.Size())
}

func TestPop(t *testing.T) {
	wg.Add(2)
	s := New[int]()
	for i := 0; i < 3000; i++ {
		s.Push(i + 30000)
	}
	go func() {
		for i := 0; i < 1300; i++ {
			fmt.Printf("11111s.Pop(): %v\n", s.Pop())
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1300; i++ {
			fmt.Printf("222222s.Pop(): %v\n", s.Pop())
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("s.Size(): %v\n", s.Size())
}

func TestPushMore(t *testing.T) {
	s := New[int]()
	s.PushMore(1, 2, 3, 4, 8, 99, -1)
	for it := s.Begin(); it != s.End(); it = it.Next() {
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

func BenchmarkPush(b *testing.B) {
	s := New[int]()
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		s.Push(rand.Intn(5000))
	}
}

func BenchmarkPushMore(b *testing.B) {
	s := New[int]()
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		s.PushMore(rand.Intn(5000), rand.Intn(5000), rand.Intn(5000), rand.Intn(5000), rand.Intn(5000), rand.Intn(5000))
	}
}

func BenchmarkPop(b *testing.B) {
	s := New[int]()
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		s.Push(rand.Intn(5000))
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
