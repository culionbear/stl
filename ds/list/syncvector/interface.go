package vector

import "stl/ds/list/list"

// Push Push value into the vector
func (m *Manager[V]) Push(v V) {
	m.mutex.Lock()
	defer func() {
		m.size++
		m.mutex.Unlock()
	}()
	if m.size < m.cap {
		m.data[m.size] = v
		return
	}
	m.cap *= 2
	tmp := make([]V, m.cap, m.cap)
	copy(tmp, m.data)
	m.data = tmp
	m.data[m.size] = v
}

// PushMore push values into the vector
func (m *Manager[V]) PushMore(values ...V) {
	for _, value := range values {
		m.Push(value)
	}
}

// Begin Get the beginning iterator of the vector
func (m *Manager[V]) Begin() list.Iterator[V] {
	return &iterator[V]{
		pointer: &m.data,
		cursor:  0,
		size:    m.size,
	}
}

// End Get the end iterator of the vector
func (m *Manager[V]) End() list.Iterator[V] {
	return nil
}

// Size Get the size of the vector
func (m *Manager[V]) Size() int {
	return m.size
}

// First Get the first value of the vector
func (m *Manager[V]) First() V {
	if m.size == 0 {
		return *new(V)
	}
	return m.data[0]
}
