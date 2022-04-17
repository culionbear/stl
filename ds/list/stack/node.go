package stack

import (
	"reflect"
	"stl/ds/list/list"
)

type node[V comparable] struct {
	next  *node[V]
	value V
}

func newNode[V comparable](v V) *node[V] {
	return &node[V]{
		next:  nil,
		value: v,
	}
}

func (m *node[V]) Value() V {
	return m.value
}

func (m *node[V]) Next() list.Iterator[V] {
	return m.next
}

func (m *node[V]) End() bool {
	return !reflect.ValueOf(m).IsNil()
}
