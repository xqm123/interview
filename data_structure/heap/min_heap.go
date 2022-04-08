package heap

import (
	"errors"
	"sync"
)

//最小堆实现
type minHeap struct {
	mu sync.RWMutex

	data  []int64 //数据
	n     uint64  //容量
	count uint64  //当前数据量
}

func NewMinHeap(n uint64) *minHeap {
	return &minHeap{
		data: make([]int64, n+1, n+1),
		n:    n,
	}
}

func (h *minHeap) AddElement(data int64) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.count >= h.n {
		return errors.New("count to reach maximum")
	}

	h.data[h.count+1] = data
	h.count++

	i := h.count
	p := i / 2
	for p > 0 && h.data[i] < h.data[p] {
		h.data[i], h.data[p] = h.data[p], h.data[i]
		i = p
		p = i / 2
	}

	return nil
}

func (h *minHeap) DelElement(index uint64) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if index == 0 || h.count == 0 || index > h.count {
		return errors.New("index error")
	}

	h.data[index] = h.data[h.count]
	h.count--

	return h.top_build_heap(index)
}

func (h *minHeap) RefBuildHeap() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	for i := (h.count / 2); i >= 1; i-- {
		if err := h.top_build_heap(i); err != nil {
			return err
		}
	}

	return nil
}

func (h *minHeap) BuildHeapFromData(data []int64) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(data) == 0 {
		return errors.New("data is empty")
	}

	if h.n < uint64(len(data)) {
		return errors.New("data len over the storable size")
	}
	//第一个元素不设置值
	copy(h.data[1:], data)
	h.count = uint64(len(data))

	for i := (h.count / 2); i >= 1; i-- {
		if err := h.top_build_heap(i); err != nil {
			return err
		}
	}

	return nil
}

func (h *minHeap) Sort() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.count <= 1 {
		return nil
	}
	count := h.count
	defer func() {
		h.count = count
	}()

	for i := count; i >= 2; i-- {
		min := h.data[1]
		h.data[1] = h.data[i]
		h.data[i] = min
		h.count--
		if err := h.top_build_heap(1); err != nil {
			return err
		}
	}

	return nil
}

func (h *minHeap) GetData() []int64 {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if h.data != nil {
		return h.data[1:]
	}
	return []int64{}
}

func (h *minHeap) GetTopValue() (int64, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if h.GetLen() < 1 {
		return 0, errors.New("data len is lt 1")
	}
	return h.data[1], nil
}

func (h *minHeap) GetCap() uint64 {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.n
}

func (h *minHeap) GetLen() uint64 {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.count
}

func (h *minHeap) top_build_heap(index uint64) error {
	if index == 0 || h.count == 0 || index > h.count {
		return errors.New("index error")
	}

	i := index
	minPos := index

	var n uint64
	for {
		n = i * 2
		if n <= h.count && h.data[n] < h.data[i] {
			minPos = n
		}
		if (n+1) <= h.count && h.data[n+1] < h.data[minPos] {
			minPos = n + 1
		}

		if i == minPos {
			break
		}
		h.data[i], h.data[minPos] = h.data[minPos], h.data[i]

		i = minPos
	}

	return nil
}

func (h *minHeap) ReplaceTopValue(val int64) error {
	h.data[1] = val
	return h.top_build_heap(1)
}
