package vector

import "sync"

type Manager[V comparable] struct {
	// data The array to store the data
	data []V
	// size The size of the vector
	size int
	// cap The capacity of the vector
	cap int
	// mutex
	mutex sync.Mutex
}

// New Get a vector with values which user put in
// Or Get a vector with the default capacity
func New[V comparable](values ...V) *Manager[V] {
	if len(values) == 0 {
		return &Manager[V]{
			data:  make([]V, 10, 10),
			size:  0,
			cap:   10,
			mutex: sync.Mutex{},
		}
	}
	m := &Manager[V]{
		data:  make([]V, 2*len(values), 2*len(values)),
		size:  0,
		cap:   2 * len(values),
		mutex: sync.Mutex{},
	}
	m.PushMore(values...)
	return m
}
