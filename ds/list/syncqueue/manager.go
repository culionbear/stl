package queue

import "sync"

type Manager[V comparable] struct {
	// front Pointer to the head of the queue
	front *node[V]
	// tail Pointer to the end of the queue
	tail *node[V]
	// size The size of the queue
	size  int
	mutex sync.Mutex
}

// New Get a new queue with values which user put in
func New[V comparable](values ...V) *Manager[V] {
	m := &Manager[V]{
		front: nil,
		tail:  nil,
		size:  0,
		mutex: sync.Mutex{},
	}
	m.PushMore(values...)
	return m
}
