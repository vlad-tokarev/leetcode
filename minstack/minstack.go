package minstack

type MinStack struct {
	heap     *FastRemovalHeap
	stack    []int
	stackRef []int
}

func Constructor() MinStack {
	return MinStack{
		heap: &FastRemovalHeap{
			data:       make([]int, 0),
			refToIndex: make(map[int]int),
			indexToRef: make(map[int]int),
		},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	ref := this.heap.Insert(val)
	this.stackRef = append(this.stackRef, ref)

}

func (this *MinStack) Pop() {
	ref := this.stackRef[len(this.stackRef)-1]
	this.stackRef = this.stackRef[:len(this.stackRef)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.heap.Delete(ref)
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.heap.GetMin()
}

// FastRemovalHeap is like regular binary heap, but allows to delete for O(log n) not only root node
// But also any element by value.
// It leverages hashtable for keeping references between values and indexes in underlying heap array
// for fast search heap array indexes
type FastRemovalHeap struct {
	data []int
	// mapping keeps references value -> indexes in underlying binary heap array
	// For fast search
	refToIndex map[int]int
	indexToRef map[int]int
}

func (h *FastRemovalHeap) swap(i, j int) {

	refI, refJ := h.indexToRef[i], h.indexToRef[j]
	h.refToIndex[refI] = j
	h.refToIndex[refJ] = i

	h.indexToRef[i] = refJ
	h.indexToRef[j] = refI

	// swap in heap array
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *FastRemovalHeap) addMapping(ref, i int) {
	h.refToIndex[ref] = i
	h.indexToRef[i] = ref
}

func (h *FastRemovalHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *FastRemovalHeap) children(i int) (int, int) {
	j := 2*i + 1
	k := 2*i + 2

	if j >= len(h.data) {
		return -1, -1
	}
	if k >= len(h.data) {
		return j, -1
	}
	return j, k
}

func (h *FastRemovalHeap) minChild(i int) (int, int) {
	i, j := h.children(i)
	switch {
	case i == -1 && j == -1:
		return -1, 0
	case i == -1:
		return j, h.data[j]
	case j == -1:
		return i, h.data[i]
	case h.data[j] < h.data[i]:
		return j, h.data[j]
	default:
		return i, h.data[i]
	}

}

func (h *FastRemovalHeap) removeRef(ref int) {
	index := h.refToIndex[ref]
	delete(h.refToIndex, ref)
	delete(h.indexToRef, index)
}

// Insert inserts item into heap, and returns index in underlying array
func (h *FastRemovalHeap) Insert(val int) int {
	h.data = append(h.data, val)
	ref := len(h.data) - 1
	i := ref
	h.addMapping(ref, i)

L:
	if i == 0 { // edge case
		return ref
	}
	val = h.data[i]
	parent := h.parent(i)
	if h.data[parent] <= val {
		return ref
	}

	h.swap(i, parent)
	i = parent
	goto L

}

func (h *FastRemovalHeap) Delete(ref int) {

	// let's find index in array of value
	// if there are multiple values in heap, take any (last)
	index := h.refToIndex[ref]

	if index == len(h.data)-1 {
		// edge case
		h.data = h.data[:len(h.data)-1]
		h.removeRef(ref)
		return
	}

	// now work with underlying array as usual
	h.swap(index, len(h.data)-1)
	h.data = h.data[:len(h.data)-1]

	val := h.data[index]
	var i, v int
L:

	i, v = h.minChild(index)
	if i == -1 { // edge case, no children
		h.removeRef(ref)
		return
	}

	if val <= v {
		// stop
		h.removeRef(ref)
		return
	}

	// swap
	h.swap(index, i)
	index = i
	goto L
}

func (h *FastRemovalHeap) GetMin() int {
	if len(h.data) == 0 {
		return 0
	}
	return h.data[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
