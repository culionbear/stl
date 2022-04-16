package stack

import "stl/ds/list/list"

// Push Push value into stack
func (m *Manager[V]) Push(v V) {
	newNode := newNode(v)
	defer func() {
		m.size++
	}()
	if m.top == nil {
		m.top = newNode
		return
	}
	newNode.next = m.top
	m.top = newNode
}

// PushMore Push values into stack
func (m *Manager[V]) PushMore(values ...V) {
	for _, value := range values {
		m.Push(value)
	}
}

// Begin Get the Beginning of the iterator
func (m *Manager[V]) Begin() list.Iterator[V] {
	return m.top
}

// End Get the end of the iterator
func (m *Manager[V]) End() list.Iterator[V] {
	return nil
}

// Size Get the size of the stack
func (m *Manager[V]) Size() int {
	return m.size
}

// First Get the value of the stack on the top
func (m *Manager[V]) First() V {
	if m.top == nil {
		return *new(V)
	}
	return m.top.value
}
