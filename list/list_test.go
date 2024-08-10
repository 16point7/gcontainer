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
