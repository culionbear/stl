package list

type List[V comparable] interface {
	Push(V)
	PushMore(...V)
	Begin() Iterator[V]
	End() Iterator[V]
	Size() int
	First() V
}