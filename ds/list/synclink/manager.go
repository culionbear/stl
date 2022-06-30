package link

import "sync"

type Manager[V comparable] struct {
	head  *node[V]
	tail  *node[V]
	size  int
	mutex sync.Mutex
}

func New[V comparable](values ...V) *Manager[V] {
	m := &Manager[V]{
		head:  nil,
		tail:  nil,
		size:  0,
		mutex: sync.Mutex{},
	}
	m.PushMore(values...)
	return m
}
