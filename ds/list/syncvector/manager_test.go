package vector

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestIterator(t *testing.T) {
	v := New[int]()
	v.PushMore(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for it := v.Begin(); it != v.End(); it = it.Next() {
		fmt.Println(it.Value())
	}
}

func TestPush(t *testing.T) {
	v := New[int]()
	wg.Add(2)
	go func() {
		for i := 0; i < 3000; i++ {
			v.Push(i + 10000)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 3000; i++ {
			v.Push(i + 30000)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("v.Size(): %v\n", v.Size())
}

func TestPopBack(t *testing.T) {
	v := New[int]()
	for i := 0; i < 3000; i++ {
		v.Push(i + 10000)
	}
	wg.Add(2)
	go func() {
		for i := 0; i < 300; i++ {
			v.PopBack()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 300; i++ {
			v.PopBack()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("v.Size(): %v\n", v.Size())
}

func TestAt(t *testing.T) {
	v := New[int]()
	for i := 0; i < 3000; i++ {
		v.Push(i + 10000)
	}
	wg.Add(2)
	go func() {
		for i := 0; i < 300; i++ {
			fmt.Printf("v.At(i): %v\n", v.At(i+10))
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 300; i++ {
			fmt.Printf("v.At(i): %v\n", v.At(i+1000))
		}
		wg.Done()
	}()
	wg.Wait()
}

func TestRemove(t *testing.T) {
	v := New[int]()
	for i := 0; i < 3000; i++ {
		v.Push(i + 10000)
	}
	wg.Add(2)
	go func() {
		for i := 0; i < 300; i++ {
			v.Remove(i + 100)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 300; i++ {
			v.Remove(i + 1000)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("v.Size(): %v\n", v.Size())
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
	v := New[int]()
	v.PushMore(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	wg.Add(2)
	go func() {
		v.Reverse()
		wg.Done()
	}()
	go func() {
		v.Reverse()
		wg.Done()
	}()
	wg.Wait()
	for it := v.Begin(); it != v.End(); it = it.Next() {
		fmt.Println(it.Value())
	}
}

func TestSort(t *testing.T) {
	v := New(1, 777, 23, 31)
	wg.Add(2)
	go func() {
		v.Sort(func(i1, i2 int) bool { return i1 < i2 })
		wg.Done()
	}()
	go func() {
		v.Sort(func(i1, i2 int) bool { return i1 < i2 })
		wg.Done()
	}()
	wg.Wait()
	for it := v.Begin(); it != v.End(); it = it.Next() {
		fmt.Print(it.Value(), " ")
	}
}

func TestInsert(t *testing.T) {
	v := New[int]()
	wg.Add(2)
	go func() {
		for i := 0; i < 3000; i++ {
			v.Insert(0, 1)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 3000; i++ {
			v.Insert(0, 1)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("v.Size(): %v\n", v.Size())
}
