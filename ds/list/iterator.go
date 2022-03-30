package list

import "stl/value"

type Iterator[V value.Object] interface {
	Value()	V
	Next() Iterator[V]
}
