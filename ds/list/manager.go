package list

import "stl/value"

type List[V value.Object] interface {
	Push(V)
	Begin() Iterator[V]
	End() Iterator[V]
	Size() int
	First() V
}