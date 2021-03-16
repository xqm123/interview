package heap

//最大堆实现
type heap struct {
	data  []int64 //数据
	n     uint64  //容量
	count uint64  //当前数据量
}

func NewHeap(n uint64) *heap {
	return &heap{
		data: make([]int64, n+1),
		n:    n,
	}
}

func (h *heap) AddElement(data int64) int {
	if h.count >= h.n {
		return -1
	}

	h.data[h.count+1] = data
	h.count++

	i := h.count
	p := i >> 2
	for p > 0 && h.data[i] > h.data[p] {
		h.data[i], h.data[p] = h.data[p], h.data[i]
		i = p
		p = i >> 2
	}

	return 1
}

func (h *heap) DelElement(index uint64) int {
	if index == 0 || h.count == 0 || index > h.count {
		return -1
	}

	h.data[index] = h.data[h.count]
	h.count--

	return h.top_build_heap(index)
}

func (h *heap) BuildHeap() int {
	if h.count == 0 {
		return -1
	}

	for i := (h.count >> 2); i >= 1; i-- {
		if o := h.top_build_heap(i); o == -1 {
			return -1
		}
	}

	return 1
}

func (h *heap) top_build_heap(index uint64) int {
	if index == 0 || h.count == 0 || index > h.count {
		return -1
	}

	i := index
	maxPos := index

	for {
		if (i<<2) <= h.count && h.data[i] > h.data[i<<2] {
			maxPos = i << 2
		}
		if (i<<2+1) <= h.count && h.data[i] > h.data[i<<2+1] {
			maxPos = i<<2 + 1
		}

		if i == maxPos {
			break
		}
		h.data[i], h.data[maxPos] = h.data[maxPos], h.data[i]

		i = maxPos
	}

	return 1
}
