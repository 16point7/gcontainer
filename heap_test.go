package heap

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	t.Run("Test Int", func(t *testing.T) {
		in := []int{5, 4, 3, 2, 1}
		less := func(a, b int) bool { return a < b }
		want := []int{1, 2, 3, 5, 4}

		got := New(in, less)

		if !reflect.DeepEqual(got.data, want) {
			t.Fatalf("Init failed. got %v, want %v", got.data, want)
		}
	})

	t.Run("Test Float", func(t *testing.T) {
		in := []float64{5.5, 4.4, 3.3, 2.2, 1.1}
		less := func(a, b float64) bool { return a < b }
		want := []float64{1.1, 2.2, 3.3, 5.5, 4.4}

		got := New(in, less)

		if !reflect.DeepEqual(got.data, want) {
			t.Fatalf("Init failed. got %v, want %v", got.data, want)
		}
	})

}

func TestPush(t *testing.T) {
	t.Run("Test Int", func(t *testing.T) {
		in := []int{5, 4, 3, 2, 1}
		less := func(a, b int) bool { return a < b }

		h := New(in, less)
		h.Push(77)

		want := []int{1, 2, 3, 5, 4, 77}

		if !reflect.DeepEqual(h.data, want) {
			t.Fatalf("Init failed. got %v, want %v", h.data, want)
		}
	})

	t.Run("Test Float", func(t *testing.T) {
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
