package heap

import (
	"errors"
	"reflect"
	"testing"
)

type person struct {
	name string
	age  uint
}

func TestInt(t *testing.T) {
	h := testNew(
		t,
		[]int{5, 4, 3, 2, 1},
		func(a, b int) bool { return a < b },
		[]int{1, 2, 3, 5, 4},
	)

	testPush(
		t,
		h,
		77,
		[]int{1, 2, 3, 5, 4, 77},
	)

	testPop(
		t,
		h,
		1,
		[]int{2, 4, 3, 5, 77},
	)
}

func TestFloat64(t *testing.T) {
	h := testNew(
		t,
		[]float64{5.5, 4.4, 3.3, 2.2, 1.1},
		func(a, b float64) bool { return a < b },
		[]float64{1.1, 2.2, 3.3, 5.5, 4.4},
	)

	testPush(
		t,
		h,
		77.77,
		[]float64{1.1, 2.2, 3.3, 5.5, 4.4, 77.77},
	)

	testPop(
		t,
		h,
		1.1,
		[]float64{2.2, 4.4, 3.3, 5.5, 77.77},
	)
}

func TestPerson(t *testing.T) {
	h := testNew(
		t,
		[]person{{name: "Andrew", age: 55}, {name: "John", age: 44}, {name: "Bryan", age: 33}, {name: "Daniel", age: 22}, {name: "Emmanuel", age: 11}},
		func(a, b person) bool { return a.age < b.age },
		[]person{{name: "Emmanuel", age: 11}, {name: "Daniel", age: 22}, {name: "Bryan", age: 33}, {name: "Andrew", age: 55}, {name: "John", age: 44}},
	)

	testPush(
		t,
		h,
		person{name: "Zachary", age: 77},
		[]person{{name: "Emmanuel", age: 11}, {name: "Daniel", age: 22}, {name: "Bryan", age: 33}, {name: "Andrew", age: 55}, {name: "John", age: 44}, {name: "Zachary", age: 77}},
	)

	testPop(
		t,
		h,
		person{name: "Emmanuel", age: 11}, []person{{name: "Daniel", age: 22}, {name: "John", age: 44}, {name: "Bryan", age: 33}, {name: "Andrew", age: 55}, {name: "Zachary", age: 77}},
	)
}

func testNew[V any](t *testing.T, in []V, less func(a, b V) bool, want []V) *heap[V] {
	t.Helper()

	h := New(in, less)

	if !reflect.DeepEqual(h.data, want) {
		t.Fatalf("New failed. got %v, want %v", h.data, want)
	}

	return h
}

func testPush[V any](t *testing.T, h *heap[V], v V, want []V) {
	t.Helper()

	h.Push(v)

	if !reflect.DeepEqual(h.data, want) {
		t.Fatalf("Push failed. got %v, want %v", h.data, want)
	}
}

func testPop[V any](t *testing.T, h *heap[V], vWant V, dataWant []V) {
	t.Helper()

	v, _ := h.Pop()
	if !reflect.DeepEqual(v, vWant) {
		t.Fatalf("Pop failed. got %v, want %v", v, vWant)
	}

	if !reflect.DeepEqual(h.data, dataWant) {
		t.Fatalf("Init failed. got %v, want %v", h.data, dataWant)
	}

	for h.Size() > 0 {
		h.Pop()
	}

	_, err := h.Pop()
	if !errors.Is(err, ErrorEmptyHeap) {
		t.Fatalf("Empty pop returned wrong error. got %v, want %v", err, ErrorEmptyHeap)
	}
}
