package stack

// Check whether the stack is empty
func (m *Manager[V]) IsEmpty() bool {
	return m.size == 0
}

// Pop the top node of the stack and return the value
func (m *Manager[V]) Pop() V {
	m.mutex.Lock()
	if m.top == nil {
		return *new(V)
	}
	defer func() {
		m.size--
	}()
	v := m.top.value
	m.top = m.top.next
	m.mutex.Unlock()
	return v
}

// Get the value of the top node
func (m *Manager[V]) Top() V {
	return m.First()
}
