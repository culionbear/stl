package link

type Manager[V comparable] struct {
	head	*node[V]
	tail	*node[V]
	size	int
}

func New[V comparable](values ...V) *Manager[V] {
	m := &Manager[V]{
		head: nil,
		tail: nil,
		size: 0,
	}
	m.PushMore(values...)
	return m
}
