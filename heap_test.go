package heap

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	t.Run("Test int", func(t *testing.T) {
		in := []int{5, 4, 3, 2, 1}
		less := func(a, b int) bool { return a < b }

		h := New(in, less)
		want := []int{1, 2, 3, 5, 4}

		if !reflect.DeepEqual(h.data, want) {
			t.Fatalf("Init failed. got %v, want %v", h.data, want)
		}
	})

	t.Run("Test float64", func(t *testing.T) {
		in := []float64{5.5, 4.4, 3.3, 2.2, 1.1}
		less := func(a, b float64) bool { return a < b }

		h := New(in, less)
		want := []float64{1.1, 2.2, 3.3, 5.5, 4.4}

		if !reflect.DeepEqual(h.data, want) {
			t.Fatalf("Init failed. got %v, want %v", h.data, want)
		}
	})

}

func TestPush(t *testing.T) {
	t.Run("Test int", func(t *testing.T) {
		in := []int{5, 4, 3, 2, 1}
		less := func(a, b int) bool { return a < b }

		h := New(in, less)
		h.Push(77)

		want := []int{1, 2, 3, 5, 4, 77}

		if !reflect.DeepEqual(h.data, want) {
			t.Fatalf("Init failed. got %v, want %v", h.data, want)
		}
	})

	t.Run("Test float64", func(t *testing.T) {
		in := []float64{5.5, 4.4, 3.3, 2.2, 1.1}
		less := func(a, b float64) bool { return a < b }

		h := New(in, less)
		h.Push(77.77)

		want := []float64{1.1, 2.2, 3.3, 5.5, 4.4, 77.77}

		if !reflect.DeepEqual(h.data, want) {
			t.Fatalf("Init failed. got %v, want %v", h.data, want)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("Test int", func(t *testing.T) {
		in := []int{5, 4, 3, 2, 1}
		less := func(a, b int) bool { return a < b }

		h := New(in, less)

		v := h.Pop()
		if v != 1 {
			t.Fatalf("Pop failed. got %v, want %v", v, 1)
		}

		data := []int{2, 4, 3, 5}
		if !reflect.DeepEqual(h.data, data) {
			t.Fatalf("Init failed. got %v, want %v", h.data, data)
		}
	})

	t.Run("Test float64", func(t *testing.T) {
		in := []float64{5.5, 4.4, 3.3, 2.2, 1.1}
		less := func(a, b float64) bool { return a < b }

		h := New(in, less)

		v := h.Pop()
		if v != 1.1 {
			t.Fatalf("Pop failed. got %v, want %v", v, 1.1)
		}

		data := []float64{2.2, 4.4, 3.3, 5.5}

		if !reflect.DeepEqual(h.data, data) {
			t.Fatalf("Init failed. got %v, want %v", h.data, data)
		}
	})

}
