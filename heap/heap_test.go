package heap

import (
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	data := []int{5, 4, 3, 2, 1}
	less := func(a, b int) bool { return a < b }

	h := New(data, less)
	want := []int{1, 2, 3, 5, 4}
	if !reflect.DeepEqual(h.data, want) {
		t.Fatalf("New failed. want %v, got %v", want, h.data)
	}

	h.Push(77)
	want = []int{1, 2, 3, 5, 4, 77}
	if !reflect.DeepEqual(h.data, want) {
		t.Fatalf("Push failed. want %v, got %v", want, h.data)
	}

	got := h.Pop()
	if got != 1 {
		t.Fatalf("non-empty Pop failed. want %v, got %v", 1, got)
	}

	for h.Size() > 0 {
		h.Pop()
	}
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("empty Pop failed. should have panicked")
		}
	}()
	h.Pop()
}
