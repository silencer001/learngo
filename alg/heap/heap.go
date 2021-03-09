package heap

type Heap struct {
	len int
	s   []int
}

func NewHeap(n int) *Heap {
	return &Heap{s: make([]int, n+1)}
}
func NewHeapCopy(a []int) *Heap {
	s := make([]int, 1)
	s = append(s, a...)
	return &Heap{
		s:   s,
		len: len(a),
	}
}
func (heap *Heap) minSon(parentIndex int) int {
	if parentIndex*2 > heap.len {
		panic("out of range")
	} else if parentIndex*2 == heap.len {
		return parentIndex * 2
	} else {
		if heap.s[parentIndex*2] >= heap.s[parentIndex*2+1] {
			return parentIndex*2 + 1
		} else {
			return parentIndex * 2
		}
	}
}
func (heap *Heap) BuildHeap() {
	for i := heap.len / 2; i > 0; i-- {
		heap.heapify(i)
	}
}
func (heap *Heap) heapify(j int) {
	for i := j; i > 0; i = i / 2 {
		minson := heap.minSon(i)
		if heap.s[i] > heap.s[minson] {
			heap.s[i], heap.s[minson] = heap.s[minson], heap.s[i]
		}
	}
}
func (heap *Heap) heapifyTop() {
	for i := 1; i <= heap.len/2; {
		minson := heap.minSon(i)
		if heap.s[i] > heap.s[minson] {
			heap.s[i], heap.s[minson] = heap.s[minson], heap.s[i]
			i = minson
		} else {
			break
		}
	}
}
func (heap *Heap) Push(v int) {
	heap.s = append(heap.s, v)
	heap.len++
	heap.heapify(heap.len / 2)
}

func (heap *Heap) PopTop() int {
	top := heap.s[1]
	heap.s[1] = heap.s[heap.len]
	heap.s = heap.s[:heap.len]
	heap.len--
	heap.heapifyTop()
	return top
}
