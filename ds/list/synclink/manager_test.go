package link

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func Print[V comparable](l *Manager[V]) {
	for it := l.Begin(); it != l.End(); it = it.Next() {
		fmt.Print(it.Value(), " ")
	}
	fmt.Println()
}

func TestIterator(t *testing.T) {
	l := New(1, 2, 3)
	Print(l)
}

func TestPush(t *testing.T) {
	l := New[int]()
	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			l.Push(i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			l.Push(i + 100)
		}
		wg.Done()
	}()
	wg.Wait()
	// Print(l)
	fmt.Printf("l.Size(): %v\n", l.Size())
}

func TestPushMore(t *testing.T) {
	l := New[int]()
	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			l.PushMore(i, 1, 2, 3, 4)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			l.PushMore(i+100, 1, 2, 3, 4)
		}
		wg.Done()
	}()
	wg.Wait()
	// Print(l)
	fmt.Printf("l.Size(): %v\n", l.Size())
}

func TestSize(t *testing.T) {
	l := New[int]()
	fmt.Println(l.Size())
	l.PushMore(1, 2, 4)
	fmt.Println(l.Size())
}

func TestFirst(t *testing.T) {
	l := New(1, 2, 4)
	fmt.Printf("l.First(): %v\n", l.First())
}

func TestAt(t *testing.T) {
	l := New(1, 2, 4, 8)
	fmt.Printf("l.At(5): %v\n", l.At(5))
	fmt.Printf("l.At(-4): %v\n", l.At(-4))
	fmt.Printf("l.At(3): %v\n", l.At(3))
	fmt.Printf("l.At(0): %v\n", l.At(0))
}

func TestSort(t *testing.T) {
	l := New(3, 2, 1, 9, 0, 6)
	wg.Add(2)
	go func() {
		l.Sort(func(i1, i2 int) bool {
			return i1 < i2
		})
		wg.Done()
	}()
	Print(l)
	go func() {
		l.Sort(func(i1, i2 int) bool {
			return i2 < i1
		})
		wg.Done()
	}()
	wg.Wait()
	Print(l)
	l.Sort(nil)
	Print(l)
}

func TestRemove(t *testing.T) {
	l := New[int]()
	for i := 0; i < 10000; i++ {
		l.Push(i)
	}
	wg.Add(5)
	go func() {
		l.Remove(88)
		wg.Done()
	}()
	go func() {
		l.Remove(88)
		wg.Done()
	}()
	go func() {
		l.Remove(88)
		wg.Done()
	}()
	go func() {
		l.Remove(88)
		wg.Done()
	}()
	go func() {
		l.Remove(88)
		wg.Done()
	}()
	wg.Wait()
	Print(l)
	fmt.Printf("l.Size(): %v\n", l.Size())
}

func TestRemoveValue(t *testing.T) {
	l := New(1, 1, 2, 4, 5, 1, 2, 7, 8, 0)
	wg.Add(2)
	go func() {
		l.RemoveValue(1)
		wg.Done()
	}()
	Print(l)
	go func() {
		l.RemoveValue(1)
		wg.Done()
	}()
	wg.Wait()
	Print(l)
}

func TestInsert(t *testing.T) {
	l := New(1, 2, 3)
	wg.Add(2)
	go func() {
		l.Insert(1, 1000)
		wg.Done()
	}()
	Print(l)
	go func() {
		l.Insert(1, 10001)
		wg.Done()
	}()
	wg.Wait()
	Print(l)
}

func BenchmarkPush(b *testing.B) {
	l := New[int]()
	wg.Add(2)
	go func() {
		for i := 0; i < b.N; i++ {
			l.Push(1)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			l.Push(1)
		}
		wg.Done()
	}()
	wg.Wait()
}

func BenchmarkPushMore(b *testing.B) {
	l := New[int]()
	wg.Add(2)
	go func() {
		for i := 0; i < b.N; i++ {
			l.PushMore(1, 1, 1, 1, 1)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			l.PushMore(1, 1, 1, 1, 1)
		}
		wg.Done()
	}()
	wg.Wait()
}

func BenchmarkAt(b *testing.B) {
	l := New(1, 1, 2, 4, 5, 1, 2, 7, 8, 0)
	for i := 0; i < b.N; i++ {
		l.At(rand.Intn(10))
	}
}

func BenchmarkSort(b *testing.B) {
	l := New[int]()
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000000; i++ {
		l.Push(rand.Intn(1000000000))
	}
	s := time.Now().Unix()
	l.Sort(func(i1, i2 int) bool {
		return i1 < i2
	})
	e := time.Now().Unix()
	fmt.Println(e - s)
}

func BenchmarkRemove(b *testing.B) {
	l := New[int]()
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000000; i++ {
		l.Push(rand.Intn(1000000000))
	}
	for i := 0; i < b.N; i++ {
		l.Remove(9)
	}
}

func BenchmarkRemoveValue(b *testing.B) {
	l := New[int]()
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000000; i++ {
		l.Push(rand.Intn(1000000000))
	}
	for i := 0; i < b.N; i++ {
		l.Remove(l.At(0))
	}
}

func BenchmarkInsert(b *testing.B) {
	l := New[int]()
	for i := 0; i < b.N; i++ {
		l.Insert(l.Size(), 1)
	}
}
