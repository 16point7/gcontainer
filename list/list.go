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
	l.sentinel.list = l
	return l
}

func (l *List[V]) Len() int {
	return l.len
}

func (l *List[V]) PushFront(v V) *Element[V] {
	return l.InsertAfter(v, &l.sentinel)
}

func (l *List[V]) PushBack(v V) *Element[V] {
	return l.InsertAfter(v, l.sentinel.prev)
}

func (l *List[V]) Front() *Element[V] {
	if l.len == 0 {
		return nil
	}
	return l.sentinel.next
}

func (l *List[V]) Back() *Element[V] {
	if l.len == 0 {
		return nil
	}
	return l.sentinel.prev
}

func (l *List[V]) Remove(e *Element[V]) {
	if e.list != l {
		return
	}
	e.next.prev = e.prev
	e.prev.next = e.next
	e.next, e.prev, e.list = nil, nil, nil
	l.len--
}

func (l *List[V]) InsertAfter(v V, pin *Element[V]) *Element[V] {
	if pin.list != l {
		return nil
	}
	e := &Element[V]{Value: v}
	l.insertAfter(e, pin)
	return e
}

func (l *List[V]) InsertBefore(v V, pin *Element[V]) *Element[V] {
	return l.InsertAfter(v, pin.prev)
}

func (l *List[V]) insertAfter(e, pin *Element[V]) {
	e.next = pin.next
	e.next.prev = e
	e.prev = pin
	pin.next = e
	e.list = l
	l.len++
}

func (l *List[V]) MoveAfter(e, pin *Element[V]) {
	l.Remove(e)
	l.insertAfter(e, pin)
}

func (l *List[V]) MoveBefore(e, pin *Element[V]) {
	l.Remove(e)
	l.insertAfter(e, pin.prev)
}

func (l *List[V]) MoveToFront(e *Element[V]) {
	l.Remove(e)
	l.insertAfter(e, &l.sentinel)
}

func (l *List[V]) MoveToBack(e *Element[V]) {
	l.Remove(e)
	l.insertAfter(e, l.sentinel.prev)
}
