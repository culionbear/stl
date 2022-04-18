package queue

import "stl/ds/list/list"

// Push value into queue
func (m *Manager[V]) Push(v V) {
	n := newNode(v)
	defer func() {
		m.tail = n
		m.size++
	}()
	if m.front == nil {
		m.front = n
		return
	}
	m.tail.next = n
}

// Push values into queue
func (m *Manager[V]) PushMore(values ...V) {
	for _, value := range values {
		m.Push(value)
	}
}

// Get the beginning of the iterator
func (m *Manager[V]) Begin() list.Iterator[V] {
	return m.front
}

// Get the end of the iterator
func (m *Manager[V]) End() list.Iterator[V] {
	return nil
}

// Get the size of the queue
func (m *Manager[V]) Size() int {
	return m.size
}

// Get the value of the queue head node
func (m *Manager[V]) First() V {
	return m.front.value
}
