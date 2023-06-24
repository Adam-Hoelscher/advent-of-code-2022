package util

type IntHeap []int

type MinIntHeap struct {
	IntHeap
}

type MaxIntHeap struct {
	IntHeap
}

func (h IntHeap) Len() int      { return len(h) }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)   { *h = append(*h, x.(int)) }

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h MaxIntHeap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }
func (h MinIntHeap) Less(i, j int) bool { return h.IntHeap[i] < h.IntHeap[j] }
