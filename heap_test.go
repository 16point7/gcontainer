package heap

import (
	"errors"
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

	got, err := h.Pop()
	if got != 1 && err == nil {
		t.Fatalf("non-empty Pop failed. want %v, got %v, err %v", 1, got, err)
	}

	for len(h.data) > 0 {
		h.Pop()
	}
	_, err = h.Pop()
	if !errors.Is(err, ErrorEmptyHeap) {
		t.Fatalf("empty Pop failed. want %v, got %v", ErrorEmptyHeap, err)
	}
}
