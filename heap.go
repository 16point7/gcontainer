package heap

type heap[V any] struct {
	data []V
	less func(a, b V) bool
}

func New[V any](data []V, less func(a, b V) bool) *heap[V] {
	h := &heap[V]{data: data, less: less}
	for i := ((len(data) - 1) - 1) / 2; i >= 0; i-- {
		h.down(i)
	}
	return h
}

func (h *heap[V]) Push(v V) {
	h.data = append(h.data, v)
	h.up(len(h.data) - 1)
}

func (h *heap[V]) up(i int) {
	for {
		parent := (i - 1) / 2

		if parent < 0 || !h.less(h.data[i], h.data[parent]) {
			break
		}

		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *heap[V]) down(i int) {
	for {
		next := 2*i + 1

		if next >= len(h.data) || next < 0 {
			break
		}

		if right := next + 1; right < len(h.data) && h.less(h.data[right], h.data[next]) {
			next = right
		}

		if !h.less(h.data[next], h.data[i]) {
			break
		}

		h.data[i], h.data[next] = h.data[next], h.data[i]
		i = next
	}
}
