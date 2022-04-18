package queue

type Manager[V comparable] struct {
	front *node[V]
	tail  *node[V]
	size  int
}

func New[V comparable](values ...V) *Manager[V] {
	m := &Manager[V]{
		front: nil,
		tail:  nil,
		size:  0,
	}
	m.PushMore(values...)
	return m
}
