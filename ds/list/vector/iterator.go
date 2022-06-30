package vector

import "stl/ds/list/list"

type iterator[V comparable] struct {
	pointer *[]V
	cursor  int
	size    int
}

func (i *iterator[V]) Value() V {
	return (*i.pointer)[i.cursor]
}

func (i *iterator[V]) Next() list.Iterator[V] {
	if i.cursor+1 >= i.size {
		return nil
	}
	i.cursor++
	return i
}
