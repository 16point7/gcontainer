package list

import "testing"

func TestPushFront(t *testing.T) {
	l := New[int]()

	e1 := l.PushFront(1)

	if l.Len() != 1 {
		t.Fatalf("invalid list size. got %d, want %d", l.Len(), 1)
	}

	if e1.Value != 1 {
		t.Fatalf("invalid element value. got %d, want %d", e1.Value, 1)
	}

	if e1.Next() != nil {
		t.Fatalf("invalid next element. got %v, want %v", e1.Next(), nil)
	}

	if e1.Prev() != nil {
		t.Fatalf("invalid next element. got %v, want %v", e1.Prev(), nil)
	}

	e2 := l.PushFront(2)

	if l.Len() != 2 {
		t.Fatalf("invalid list size. got %d, want %d", l.Len(), 2)
	}

	if e2.Value != 2 {
		t.Fatalf("invalid element value. got %d, want %d", e2.Value, 2)
	}

	if e2.Next() != e1 {
		t.Fatalf("invalid next element. got %v, want %v", e2.Next(), e1)
	}

	if e2.Next().Next() != nil {
		t.Fatalf("invalid next next element. got %v, want %v", e2.Next().Next(), nil)
	}

	if e2.Prev() != nil {
		t.Fatalf("invalid next element. got %v, want %v", e2.Prev(), nil)
	}

	e3 := l.PushFront(3)

	if l.Len() != 3 {
		t.Fatalf("invalid list size. got %d, want %d", l.Len(), 3)
	}

	if e3.Value != 3 {
		t.Fatalf("invalid element value. got %d, want %d", e3.Value, 3)
	}

	if e3.Next() != e2 {
		t.Fatalf("invalid next element. got %v, want %v", e3.Next(), e2)
	}

	if e3.Next().Next() != e1 {
		t.Fatalf("invalid next next element. got %v, want %v", e3.Next().Next(), e1)
	}

	if e3.Next().Next().Next() != nil {
		t.Fatalf("invalid next next element. got %v, want %v", e3.Next().Next().Next(), nil)
	}

	if e3.Prev() != nil {
		t.Fatalf("invalid next element. got %v, want %v", e3.Prev(), nil)
	}
}

func TestPushBack(t *testing.T) {
	l := New[int]()

	e1 := l.PushBack(1)

	if l.Len() != 1 {
		t.Fatalf("invalid list size. got %d, want %d", l.Len(), 1)
	}

	if e1.Value != 1 {
		t.Fatalf("invalid element value. got %d, want %d", e1.Value, 1)
	}

	if e1.Next() != nil {
		t.Fatalf("invalid next element. got %v, want %v", e1.Next(), nil)
	}

	if e1.Prev() != nil {
		t.Fatalf("invalid next element. got %v, want %v", e1.Prev(), nil)
	}

	e2 := l.PushBack(2)

	if l.Len() != 2 {
		t.Fatalf("invalid list size. got %d, want %d", l.Len(), 2)
	}

	if e2.Value != 2 {
		t.Fatalf("invalid element value. got %d, want %d", e2.Value, 2)
	}

	if e2.Prev() != e1 {
		t.Fatalf("invalid next element. got %v, want %v", e2.Prev(), e1)
	}

	if e2.Prev().Prev() != nil {
		t.Fatalf("invalid next next element. got %v, want %v", e2.Prev().Prev(), nil)
	}

	if e2.Next() != nil {
		t.Fatalf("invalid next element. got %v, want %v", e2.Next(), nil)
	}

	e3 := l.PushBack(3)

	if l.Len() != 3 {
		t.Fatalf("invalid list size. got %d, want %d", l.Len(), 3)
	}

	if e3.Value != 3 {
		t.Fatalf("invalid element value. got %d, want %d", e3.Value, 3)
	}

	if e3.Prev() != e2 {
		t.Fatalf("invalid next element. got %v, want %v", e3.Prev(), e2)
	}

	if e3.Prev().Prev() != e1 {
		t.Fatalf("invalid next next element. got %v, want %v", e3.Prev().Prev(), e1)
	}

	if e3.Prev().Prev().Prev() != nil {
		t.Fatalf("invalid next next element. got %v, want %v", e3.Prev().Prev().Prev(), nil)
	}

	if e3.Next() != nil {
		t.Fatalf("invalid next element. got %v, want %v", e3.Next(), nil)
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
		t.Fatalf("Invalid front element. got %v, want %v", l.Back(), e)
	}

	l.PushFront(4)

	if l.Back() != e {
		t.Fatalf("Invalid front element. got %v, want %v", l.Back(), e)
	}
}
