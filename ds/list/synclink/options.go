package link

func (m *Manager[V]) At(i int) V {
	m.mutex.Lock()
	if abs(i) >= m.size && i > 0 {
		return *new(V)
	}
	n := m.head
	sum := i
	if i < 0 {
		sum += m.size
	}
	for j := 0; j < sum; j++ {
		n = n.next
	}
	m.mutex.Unlock()
	return n.value
}

func (m *Manager[V]) Sort(compare func(V, V) bool) {
	m.mutex.Lock()
	if m.head == nil || m.size <= 1 || compare == nil {
		return
	}
	for n := m.head; n.next != nil; n = n.next {
		target := n
		for k := n.next; k != nil; k = k.next {
			if compare(k.value, target.value) {
				target = k
			}
		}
		n.value, target.value = target.value, n.value
	}
	m.mutex.Unlock()
}

func (m *Manager[V]) Remove(i int) V {
	m.mutex.Lock()
	if abs(i) >= m.size && i > 0 {
		return *new(V)
	}
	defer func() {
		m.size--
		m.mutex.Unlock()
	}()
	sum := i
	if i < 0 {
		sum += m.size
	}
	if sum == 0 {
		v := m.head.value
		m.head = m.head.next
		return v
	}
	for i, n := 1, m.head; n.next != nil; n, i = n.next, i+1 {
		if i == sum {
			v := n.next.value
			n.next = n.next.next
			return v
		}
	}
	return *new(V)
}

func (m *Manager[V]) RemoveValue(v V) {
	m.mutex.Lock()
	for m.head != nil && m.head.value == v {
		m.head = m.head.next
		m.size--
	}
	if m.head == nil {
		return
	}
	for n := m.head; n.next != nil; {
		if n.next.value == v {
			n.next = n.next.next
			m.size--
		} else {
			n = n.next
		}
	}
	m.mutex.Unlock()
}

func (m *Manager[V]) Insert(i int, v V) {
	m.mutex.Lock()
	if abs(i) >= m.size && i > 0 {
		m.Push(v)
		return
	}
	defer func() {
		m.size++
		m.mutex.Unlock()
	}()
	sum := i
	if sum < 0 {
		sum += m.size
	}
	if sum == 0 {
		n := newNode(v)
		n.next = m.head
		m.head = n
		return
	}
	for i, n := 1, m.head.next; n != nil; i, n = i+1, n.next {
		if i == sum {
			k := newNode(n.value)
			k.next = n.next
			n.value = v
			n.next = k
			return
		}
	}
}
