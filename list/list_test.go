package list

import "testing"

func TestPushFront(t *testing.T) {
	l := New[int]()

	e1 := l.PushFront(1)

	if l.Len() != 1 {
		t.Fatalf("Invalid list size. got %d, want %d", l.Len(), 1)
	}

	if e1.Value != 1 {
		t.Fatalf("Invalid element value. got %d, want %d", e1.Value, 1)
	}

	if e1.Next() != nil {
		t.Fatalf("Invalid next element. got %v, want %v", e1.Next(), nil)
	}

	if e1.Prev() != nil {
		t.Fatalf("Invalid prev element. got %v, want %v", e1.Prev(), nil)
	}

	e2 := l.PushFront(2)

	if l.Len() != 2 {
		t.Fatalf("Invalid list size. got %d, want %d", l.Len(), 2)
	}

	if e2.Value != 2 {
		t.Fatalf("Invalid element value. got %d, want %d", e2.Value, 2)
	}

	if e2.Next() != e1 {
		t.Fatalf("Invalid next element. got %v, want %v", e2.Next(), e1)
	}

	if e2.Next().Next() != nil {
		t.Fatalf("Invalid next next element. got %v, want %v", e2.Next().Next(), nil)
	}

	if e2.Prev() != nil {
		t.Fatalf("Invalid prev element. got %v, want %v", e2.Prev(), nil)
	}

	e3 := l.PushFront(3)

	if l.Len() != 3 {
		t.Fatalf("Invalid list size. got %d, want %d", l.Len(), 3)
	}

	if e3.Value != 3 {
		t.Fatalf("Invalid element value. got %d, want %d", e3.Value, 3)
	}

	if e3.Next() != e2 {
		t.Fatalf("Invalid next element. got %v, want %v", e3.Next(), e2)
	}

	if e3.Next().Next() != e1 {
		t.Fatalf("Invalid next next element. got %v, want %v", e3.Next().Next(), e1)
	}

	if e3.Next().Next().Next() != nil {
		t.Fatalf("Invalid next next next element. got %v, want %v", e3.Next().Next().Next(), nil)
	}

	if e3.Prev() != nil {
		t.Fatalf("Invalid prev element. got %v, want %v", e3.Prev(), nil)
	}
}

func TestPushBack(t *testing.T) {
	l := New[int]()

	e1 := l.PushBack(1)

	if l.Len() != 1 {
		t.Fatalf("Invalid list size. got %d, want %d", l.Len(), 1)
	}

	if e1.Value != 1 {
		t.Fatalf("Invalid element value. got %d, want %d", e1.Value, 1)
	}

	if e1.Prev() != nil {
		t.Fatalf("Invalid prev element. got %v, want %v", e1.Prev(), nil)
	}

	if e1.Next() != nil {
		t.Fatalf("Invalid next element. got %v, want %v", e1.Next(), nil)
	}

	e2 := l.PushBack(2)

	if l.Len() != 2 {
		t.Fatalf("Invalid list size. got %d, want %d", l.Len(), 2)
	}

	if e2.Value != 2 {
		t.Fatalf("Invalid element value. got %d, want %d", e2.Value, 2)
	}

	if e2.Prev() != e1 {
		t.Fatalf("Invalid prev element. got %v, want %v", e2.Prev(), e1)
	}

	if e2.Prev().Prev() != nil {
		t.Fatalf("Invalid prev prev element. got %v, want %v", e2.Prev().Prev(), nil)
	}

	if e2.Next() != nil {
		t.Fatalf("Invalid next element. got %v, want %v", e2.Next(), nil)
	}

	e3 := l.PushBack(3)

	if l.Len() != 3 {
		t.Fatalf("Invalid list size. got %d, want %d", l.Len(), 3)
	}

	if e3.Value != 3 {
		t.Fatalf("Invalid element value. got %d, want %d", e3.Value, 3)
	}

	if e3.Prev() != e2 {
		t.Fatalf("Invalid prev element. got %v, want %v", e3.Prev(), e2)
	}

	if e3.Prev().Prev() != e1 {
		t.Fatalf("Invalid prev prev element. got %v, want %v", e3.Prev().Prev(), e1)
	}

	if e3.Prev().Prev().Prev() != nil {
		t.Fatalf("Invalid prev prev prev element. got %v, want %v", e3.Prev().Prev().Prev(), nil)
	}

	if e3.Next() != nil {
		t.Fatalf("Invalid next element. got %v, want %v", e3.Next(), nil)
	}
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

	l.Remove(e2)

	if l.Len() != 2 {
		t.Fatalf("Invalid list size. got %d, want %d", l.Len(), 2)
	}

	if l.Front() != e3 {
		t.Fatalf("Invalid front element. got %v, want %v", l.Front(), e3)
	}

	if e2.Next() != nil || e2.Prev() != nil {
		t.Fatalf("Invalid next and prev. got %v and %v, want %v and %v", e2.Next(), e2.Prev(), nil, nil)
	}

	l.Remove(e2)

	if l.Len() != 2 {
		t.Fatalf("Invalid list size. got %d, want %d", l.Len(), 2)
	}

	l.Remove(e3)

	if l.Len() != 1 {
		t.Fatalf("Invalid list size. got %d, want %d", l.Len(), 1)
	}

	if l.Front() != e1 {
		t.Fatalf("Invalid front element. got %v, want %v", l.Front(), e1)
	}

	l.Remove(e1)

	if l.Len() != 0 {
		t.Fatalf("Invalid list size. got %d, want %d", l.Len(), 0)
	}

	if l.Front() != nil {
		t.Fatalf("Invalid front element. got %v, want %v", l.Front(), nil)
	}

	l2 := New[int]()
	l2.PushFront(1)
	l2.PushFront(2)
	e4 := l2.PushFront(3)

	e5 := &Element[int]{Value: 3}
	l2.Remove(e5)

	if l2.Len() != 3 {
		t.Fatalf("Invalid list size. got %d, want %d", l2.Len(), 3)
	}

	if l2.Front() != e4 {
		t.Fatalf("Invalid front element. got %v, want %v", l2.Front(), e4)
	}
}
