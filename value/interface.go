package value

type Object interface {
	comparable
	HashCode()	uint32
}
