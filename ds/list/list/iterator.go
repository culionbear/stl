package list

type Iterator[V comparable] interface {
	Value() V
	Next() Iterator[V]
	End() bool
}
