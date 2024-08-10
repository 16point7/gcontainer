package list

type Element struct {
	Value      int
	next, prev *Element
	list       *List
}

func (e *Element) Next() *Element {
	if next := e.next; e.list != nil && &e.list.sentinel != next {
		return next
	}
	return nil
}

func (e *Element) Prev() *Element {
	if prev := e.prev; e.list != nil && &e.list.sentinel != prev {
		return prev
	}
	return nil
}

type List struct {
	len      int
	sentinel Element
}

func New() *List {
	l := &List{}
	l.sentinel.next = &l.sentinel
	l.sentinel.prev = &l.sentinel
	return l
}

func (l *List) Len() int {
	return l.len
}

func (l *List) PushFront(v int) *Element {
	e := &Element{Value: v, list: l}
	e.next = l.sentinel.next
	l.sentinel.next.prev = e
	e.prev = &l.sentinel
	l.sentinel.next = e
	l.len++
	return e
}
