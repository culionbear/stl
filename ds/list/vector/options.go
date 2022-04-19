package vector

// At Get the ith value of the vector
// If the i < 0, return value of the (i + m.size)th last
func (m *Manager[V]) At(i int) V {
	if abs(i) >= m.size && i > 0 {
		return *new(V)
	}
	if i < 0 {
		i += m.size
	}
	return m.data[i]
}

// Front Get the first elem of the vector
func (m *Manager[V]) Front() V {
	return m.First()
}

// Back Get the last elem of the vector
func (m *Manager[V]) Back() V {
	if m.size == 0 {
		return *new(V)
	}
	return m.data[m.size-1]
}

// IsEmpty Check the vector is empty
func (m *Manager[V]) IsEmpty() bool {
	return m.size == 0
}

// Clear Clear the vector
func (m *Manager[V]) Clear() {
	m.data = make([]V, 10, 10)
	m.size = 0
	m.cap = 10
}

// PushBack Push value into the vector
func (m *Manager[V]) PushBack(v V) {
	m.Push(v)
}

// PopBack Pop value from the end of the vector and get the value
func (m *Manager[V]) PopBack() V {
	if m.size <= 0 {
		return *new(V)
	}
	defer func() {
		m.size--
		if m.size < m.cap/2 {
			m.cap = m.cap / 2
			tmp := make([]V, m.cap, m.cap)
			copy(tmp, m.data)
			m.data = tmp
		}
	}()
	v := m.data[m.size-1]
	return v
}

// Remove Remove the ith elem of the vector
// If i < 0, return the value of the ith last
func (m *Manager[V]) Remove(i int) V {
	if abs(i) >= m.size && i > 0 {
		return *new(V)
	}
	defer func() {
		m.size--

	}()
	if i < 0 {
		i += m.size
		if m.size < m.cap/2 {
			m.cap = m.cap / 2
			tmp := make([]V, m.cap, m.cap)
			copy(tmp, m.data)
			m.data = tmp
		}
	}
	v := m.data[i]
	m.data = append(m.data[:i], m.data[i+1:]...)
	return v
}

// Insert Insert value into the vector at the ith place
func (m *Manager[V]) Insert(i int, v V) {
	if abs(i) >= m.size && i > 0 {
		m.Push(v)
		return
	}
	defer func() {
		m.size++
	}()
	if i < 0 {
		i += m.size
	}
	if m.size >= m.cap {
		tmp := make([]V, m.cap, m.cap)
		copy(tmp, m.data)
		m.data = tmp
	}
	var k int
	for k = m.size; k > 0 && k > i; k-- {
		m.data[k] = m.data[k-1]
	}
	m.data[k] = v
}

// Sort Sort the vector with the compare function
func (m *Manager[V]) Sort(compare func(V, V) bool) {
	if m.size <= 1 || compare == nil {
		return
	}
	if m.size < m.cap {
		// Free unused space
		tmp := make([]V, m.size, m.size)
		copy(tmp, m.data)
		m.data = tmp
		m.cap = m.size
	}
	for i := 0; i < m.size; i++ {
		for j := i + 1; j < m.size; j++ {
			if compare(m.data[j], m.data[i]) {
				m.data[i], m.data[j] = m.data[j], m.data[i]
			}
		}
	}
}

// Reverse Reverse the vector
func (m *Manager[V]) Reverse() {
	if m.size < m.cap {
		// Free unused space
		tmp := make([]V, m.size, m.size)
		copy(tmp, m.data)
		m.data = tmp
		m.cap = m.size
	}
	for i := 0; i < m.size/2; i++ {
		m.data[i], m.data[m.size-i-1] = m.data[m.size-i-1], m.data[i]
	}
}
