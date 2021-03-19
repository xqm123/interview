package heap

//最大堆实现
type maxHeap struct {
	data  []int64 //数据
	n     uint64  //容量
	count uint64  //当前数据量
}

func NewMaxHeap(n uint64) *maxHeap {
	return &maxHeap{
		data: make([]int64, n+1),
		n:    n,
	}
}

//往堆添加元素，每一次添加需从下往上堆化 O(log(n))
func (h *maxHeap) AddElement(data int64) int {
	if h.count >= h.n {
		return -1
	}

	h.data[h.count+1] = data
	h.count++

	i := h.count
	p := i / 2
	for p > 0 && h.data[i] > h.data[p] {
		h.data[i], h.data[p] = h.data[p], h.data[i]
		i = p
		p = i / 2
	}

	return 1
}

//删除元素 O(log(n))
func (h *maxHeap) DelElement(index uint64) int {
	if index == 0 || h.count == 0 || index > h.count {
		return -1
	}

	h.data[index] = h.data[h.count]
	h.count--

	return h.top_build_heap(index)
}

//重新建堆 O(n)
func (h *maxHeap) RefBuildHeap() int {

	for i := (h.count / 2); i >= 1; i-- {
		if o := h.top_build_heap(i); o == -1 {
			return -1
		}
	}

	return 1
}

//按照来源数据建堆 O(n)
func (h *maxHeap) BuildHeapFromData(data []int64) int {
	if len(data) == 0 {
		return -1
	}

	if h.n < uint64(len(data)) {
		return -1
	}
	//第一个元素不设置值
	copy(h.data[1:], data)
	h.count = uint64(len(data))

	if o := h.RefBuildHeap(); o == -1 {
		return -1
	}

	return 1
}

//最大堆 从小到大排序 O(nlog(n))
func (h *maxHeap) Sort() int {
	if h.count <= 1 {
		return 1
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
		if o := h.top_build_heap(1); o == -1 {
			return -1
		}
	}

	return 1
}

func (h *maxHeap) GetData() []int64 {
	if h.data != nil {
		return h.data[1:]
	}
	return []int64{}
}

func (h *maxHeap) GetCap() uint64 {
	return h.n
}

func (h *maxHeap) GetLen() uint64 {
	return h.count
}

func (h *maxHeap) top_build_heap(index uint64) int {
	if index == 0 || h.count == 0 || index > h.count {
		return -1
	}

	i := index
	maxPos := index

	var n uint64
	for {
		n = i * 2
		if n <= h.count && h.data[n] > h.data[i] {
			maxPos = n
		}
		if (n+1) <= h.count && h.data[n+1] > h.data[maxPos] {
			maxPos = n + 1
		}

		if i == maxPos {
			break
		}
		h.data[i], h.data[maxPos] = h.data[maxPos], h.data[i]

		i = maxPos
	}

	return 1
}
