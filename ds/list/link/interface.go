package link

import "stl/ds/list/list"

func (m *Manager[V]) Push(v V) {
	n := newNode(v)
	defer func() {
		m.tail = n
		m.size ++
	}()
	if m.head == nil {
		m.head = n
		return
	}
	m.tail.next = n
}

func (m *Manager[V]) PushMore(list ...V) {
	for _, v := range list {
		n := newNode(v)
		if m.head == nil {
			m.head = n
			m.tail = n
		} else {
			m.tail.next = n
			m.tail = n
		}
	}
	m.size += len(list)
}

func (m *Manager[V]) Begin() list.Iterator[V] {
	return m.head
}

func (m *Manager[V]) End() list.Iterator[V] {
	return nil
}

func (m *Manager[V]) Size() int {
	return m.size
}

func (m *Manager[V]) First() V {
	if m.head == nil {
		return *new(V)
	}
	return m.head.value
}
