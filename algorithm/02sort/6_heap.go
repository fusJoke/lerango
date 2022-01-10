package main

type Heap struct {
	Size int
	Array []int
}

func NewHeap(array []int) *Heap {
	heap := new(Heap)
	heap.Array = array
	return heap
}

func(h *Heap) Push(x int) {
	if h.Size == 0 {
		h.Array[0] = x
		h.Size++
		return
	}

	i := h.Size

	for i > 0 {
		parent := (i -1) /2
		if x <= h.Array[parent] {
			break
		}

		h.Array[i] = h.Array[parent]
		i = parent
	}
	h.Array[i] = x

	h.Size++
}

func Pop(h *Heap) int {
	if h.Size == 0 {
		return -1
	}
	ret := h.Array[0]

	h.Size--
	x := h.Array[h.Size]
	h.Array[h.Size] = ret

	i := 0
	for {
		a := 2*i + 1
		b := 2*i + 2

		if a >= h.Size {
			break
		}

		if b < h.Size && h.Array[b] > h.Array[a] {
			a = b
		}

		if x >= h.Array[a] {
			break
		}

		h.Array[i] = h.Array[a]

		i = a
	}
	h.Array[i] = x
	return ret
}


