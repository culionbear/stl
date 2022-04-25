package stack

import "sync"

type Manager[V comparable] struct {
	// top Pointer to the top of the stack
	top *node[V]
	// size The size of the stack
	size  int
	mutex sync.Mutex
}

// Get a new stack with values which user put in
func New[V comparable](values ...V) *Manager[V] {
	m := &Manager[V]{
		top:   nil,
		size:  0,
		mutex: sync.Mutex{},
	}
	m.PushMore(values...)
	return m
}
