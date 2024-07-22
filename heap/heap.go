package heap

type heap[V any] struct {
	data []V
	less func(V, V) bool
}

func New[V any](data []V, less func(V, V) bool) *heap[V] {
	h := &heap[V]{data: data, less: less}
	for i := h.Size()/2 - 1; i >= 0; i-- {
		h.down(i)
	}
	return h
}

func (h *heap[V]) Push(v V) {
	h.data = append(h.data, v)
	h.up(h.Size() - 1)
}

func (h *heap[V]) Pop() V {
	if h.Size() <= 0 {
		panic("cannot pop empty heap")
	}
	res := h.data[0]
	n := h.Size() - 1
	h.data[0] = h.data[n]
	h.data = h.data[:n]
	h.down(0)
	return res
}

func (h *heap[V]) Size() int {
	return len(h.data)
}

func (h *heap[V]) up(i int) {
	for {
		parent := (i - 1) / 2

		if i == parent || !h.less(h.data[i], h.data[parent]) {
			break
		}

		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *heap[V]) down(i int) {
	for {
		next := 2*i + 1

		if next >= h.Size() || next < 0 {
			break
		}

		if right := next + 1; right < h.Size() && h.less(h.data[right], h.data[next]) {
			next = right
		}

		if !h.less(h.data[next], h.data[i]) {
			break
		}

		h.data[i], h.data[next] = h.data[next], h.data[i]
		i = next
	}
}
