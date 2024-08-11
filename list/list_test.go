package list

import (
	"testing"
)

func TestPushFront(t *testing.T) {
	l := New[int]()

	e1 := l.PushFront(1)

	if e1.Value != 1 {
		t.Fatalf("Invalid element value. got %d, want %d", e1.Value, 1)
	}

	validateListOrdering(t, l, []*Element[int]{e1})

	e2 := l.PushFront(2)

	if e2.Value != 2 {
		t.Fatalf("Invalid element value. got %d, want %d", e2.Value, 2)
	}

	validateListOrdering(t, l, []*Element[int]{e2, e1})

	e3 := l.PushFront(3)

	if e3.Value != 3 {
		t.Fatalf("Invalid element value. got %d, want %d", e3.Value, 3)
	}

	validateListOrdering(t, l, []*Element[int]{e3, e2, e1})
}

func TestPushBack(t *testing.T) {
	l := New[int]()

	e1 := l.PushBack(1)

	if e1.Value != 1 {
		t.Fatalf("Invalid element value. got %d, want %d", e1.Value, 1)
	}

	validateListOrdering(t, l, []*Element[int]{e1})

	e2 := l.PushBack(2)

	if e2.Value != 2 {
		t.Fatalf("Invalid element value. got %d, want %d", e2.Value, 2)
	}

	validateListOrdering(t, l, []*Element[int]{e1, e2})

	e3 := l.PushBack(3)

	if e3.Value != 3 {
		t.Fatalf("Invalid element value. got %d, want %d", e3.Value, 3)
	}

	validateListOrdering(t, l, []*Element[int]{e1, e2, e3})
}

func TestFront(t *testing.T) {
	l := New[int]()

	l.PushFront(1)
	l.PushFront(2)
	e := l.PushFront(3)

	if l.Front() != e {
		t.Fatalf("Invalid front element. got %v, want %v", l.Front(), e)
	}

	l.PushBack(4)

	if l.Front() != e {
		t.Fatalf("Invalid front element. got %v, want %v", l.Front(), e)
	}
}

func TestBack(t *testing.T) {
	l := New[int]()

	l.PushBack(1)
	l.PushBack(2)
	e := l.PushBack(3)

	if l.Back() != e {
		t.Fatalf("Invalid back element. got %v, want %v", l.Back(), e)
	}

	l.PushFront(4)

	if l.Back() != e {
		t.Fatalf("Invalid back element. got %v, want %v", l.Back(), e)
	}
}

func TestRemove(t *testing.T) {
	l := New[int]()

	e1 := l.PushFront(1)
	e2 := l.PushFront(2)
	e3 := l.PushFront(3)

	validateListOrdering(t, l, []*Element[int]{e3, e2, e1})

	l.Remove(e2)

	validateListOrdering(t, l, []*Element[int]{e3, e1})

	l.Remove(e2)

	validateListOrdering(t, l, []*Element[int]{e3, e1})

	l.Remove(e3)

	validateListOrdering(t, l, []*Element[int]{e1})

	l.Remove(e1)

	validateListOrdering(t, l, []*Element[int]{})

	l.Remove(e1)

	validateListOrdering(t, l, []*Element[int]{})

	e4 := l.PushFront(1)
	e5 := l.PushFront(2)
	e6 := l.PushFront(3)

	validateListOrdering(t, l, []*Element[int]{e6, e5, e4})

	e7 := &Element[int]{Value: 3}
	l.Remove(e7)

	validateListOrdering(t, l, []*Element[int]{e6, e5, e4})
}

func TestInsertAfter(t *testing.T) {
	l := New[int]()

	e1 := l.PushFront(1)

	validateListOrdering(t, l, []*Element[int]{e1})

	e2 := l.InsertAfter(2, e1)

	validateListOrdering(t, l, []*Element[int]{e1, e2})

	e3 := l.InsertAfter(4, e2)

	validateListOrdering(t, l, []*Element[int]{e1, e2, e3})

	e4 := l.InsertAfter(5, e1)

	validateListOrdering(t, l, []*Element[int]{e1, e4, e2, e3})
}

func TestInsertBefore(t *testing.T) {
	l := New[int]()

	e1 := l.PushFront(1)

	validateListOrdering(t, l, []*Element[int]{e1})

	e2 := l.InsertBefore(2, e1)

	validateListOrdering(t, l, []*Element[int]{e2, e1})

	e3 := l.InsertBefore(3, e2)

	validateListOrdering(t, l, []*Element[int]{e3, e2, e1})

	e4 := l.InsertBefore(3, e1)

	validateListOrdering(t, l, []*Element[int]{e3, e2, e4, e1})
}

func TestMoveAfter(t *testing.T) {
	l := New[int]()

	e1 := l.PushFront(1)
	e2 := l.InsertAfter(2, e1)

	validateListOrdering(t, l, []*Element[int]{e1, e2})

	l.MoveAfter(e1, e2)

	validateListOrdering(t, l, []*Element[int]{e2, e1})

	e3 := l.PushBack(3)

	validateListOrdering(t, l, []*Element[int]{e2, e1, e3})

	l.MoveAfter(e3, e2)

	validateListOrdering(t, l, []*Element[int]{e2, e3, e1})

	l.MoveAfter(e3, e1)

	validateListOrdering(t, l, []*Element[int]{e2, e1, e3})

	l.MoveAfter(e2, e1)

	validateListOrdering(t, l, []*Element[int]{e1, e2, e3})

	l.MoveAfter(e1, e3)

	validateListOrdering(t, l, []*Element[int]{e2, e3, e1})
}

func TestMoveBefore(t *testing.T) {
	l := New[int]()

	e1 := l.PushFront(1)
	e2 := l.InsertAfter(2, e1)

	validateListOrdering(t, l, []*Element[int]{e1, e2})

	l.MoveBefore(e2, e1)

	validateListOrdering(t, l, []*Element[int]{e2, e1})

	e3 := l.PushBack(3)

	validateListOrdering(t, l, []*Element[int]{e2, e1, e3})

	l.MoveBefore(e3, e1)

	validateListOrdering(t, l, []*Element[int]{e2, e3, e1})

	l.MoveBefore(e1, e2)

	validateListOrdering(t, l, []*Element[int]{e1, e2, e3})

	l.MoveBefore(e2, e1)

	validateListOrdering(t, l, []*Element[int]{e2, e1, e3})
}

func validateListOrdering[V any](t *testing.T, l *List[V], want []*Element[V]) {
	t.Helper()

	if l.Len() != len(want) {
		t.Fatalf("Invalid list length. got %d, want %d", l.Len(), len(want))
	}

	f, b := l.Front(), l.Back()
	for i := 0; i < len(want); i++ {
		if wantF, wantB := want[i], want[len(want)-1-i]; f != wantF || b != wantB {
			t.Fatalf("Invalid element ordering. got f: %v and b: %v, want f: %v and b: %v", f, b, wantF, wantB)
		}
		f, b = f.Next(), b.Prev()
	}

	if f != nil || b != nil {
		t.Fatalf("Invalid terminal elements. got f: %v and b: %v, want f: %v and b: %v", f, b, nil, nil)
	}
}
