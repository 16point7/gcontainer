package list

type Element[V any] struct {
	Value      V
	next, prev *Element[V]
	list       *List[V]
}

func (e *Element[V]) Next() *Element[V] {
	if next := e.next; e.list != nil && &e.list.sentinel != next {
		return next
	}
	return nil
}

func (e *Element[V]) Prev() *Element[V] {
	if prev := e.prev; e.list != nil && &e.list.sentinel != prev {
		return prev
	}
	return nil
}

type List[V any] struct {
	len      int
	sentinel Element[V]
}

func New[V any]() *List[V] {
	l := &List[V]{}
	l.sentinel.next = &l.sentinel
	l.sentinel.prev = &l.sentinel
	return l
}

func (l *List[V]) Len() int {
	return l.len
}

func (l *List[V]) PushFront(v V) *Element[V] {
	e := &Element[V]{Value: v, list: l}
	e.next = l.sentinel.next
	l.sentinel.next.prev = e
	e.prev = &l.sentinel
	l.sentinel.next = e
	l.len++
	return e
}

func (l *List[V]) PushBack(v V) *Element[V] {
	e := &Element[V]{Value: v, list: l}
	e.prev = l.sentinel.prev
	l.sentinel.prev.next = e
	e.next = &l.sentinel
	l.sentinel.prev = e
	l.len++
	return e
}
