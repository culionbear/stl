package queue

// Check whether the queue is empty
func (m *Manager[V]) IsEmpty() bool {
	return m.size == 0
}

// Pop the first node and return the value of the first node
func (m *Manager[V]) Pop() V {
	if m.front == nil {
		return *new(V)
	}
	defer func() {
		m.size--
	}()
	v := m.front.value
	m.front = m.front.next
	return v
}

// Get the value of the first node
func (m *Manager[V]) Front() V {
	return m.First()
}

// Get the value of the last node
func (m *Manager[V]) Tail() V {
	return m.tail.value
}
