package heap

import "log"

type minHeap struct {
	heap
}

// Create an empty minHeap
func NewMinHeap() *minHeap {
	h := minHeap{heap{}}
	return &h
}

// Set nodes-field of minHeap: create node slice from interface{} slice via InitNodes(...)
func (h *minHeap) SetNodes(interSl []interface{}, scoreFunc func(interface{}) int) {
	h.nodes = InitNodes(interSl, scoreFunc)
	l := len(h.nodes)

	for i := l / 2; i >= 0; i-- {
		h.HeapifyDown(i)
	}
}

// Insert node at last position, re-heapify
func (h *minHeap) insertNode(n node) {
	h.nodes = append(h.nodes, n)
	h.HeapifyUp(len(h.nodes) - 1)
}

// Insert node at last position, re-heapify
func (h *minHeap) Insert(key interface{}, score int) {
	h.insertNode(node{key: key, score: score})
}

// Delete top node, fill blank with last node and re-heapify
func (h *minHeap) DeleteTop() {
	l := len(h.nodes) - 1
	if l < 0 {
		log.Fatal("Can't delete Top node. Heap is already empty...")
	}

	h.nodes[0], h.nodes[l] = h.nodes[l], h.nodes[0]
	h.nodes = h.nodes[:l]
	h.HeapifyDown(0)
}

// Extract top node, fill blank with last node and re-heapify
func (h *minHeap) extractTopNode() node {
	l := len(h.nodes) - 1
	if l < 0 {
		log.Fatal("Can't extract from empty heap...")
	}

	ex := h.nodes[0]
	h.nodes[0], h.nodes[l] = h.nodes[l], h.nodes[0]
	h.nodes = h.nodes[:l]
	h.HeapifyDown(0)

	return ex
}

// Extract top node via extractTopNode(...) and return its key and score.
func (h *minHeap) ExtractTop() (interface{}, int) {
	exNode := h.extractTopNode()
	return exNode.key, exNode.score
}

// Return top n node.keys and associated node.scores in two sorted slices
func (h *minHeap) Peek(n int) ([]interface{}, []int) {
	out := []node{}
	var (
		k1, k2   int  // index of [left:1|right:2] child node
		sc1, sc2 int  // score of [left:1|right:2] child node
		topIdx   int  // index of current top node of hh as indexed in original heap
		ok1, ok2 bool // ok if children exist
	)

	// hh (heap of heaps) stores nodes which contain information about the original heap:
	//   key: piTup{pointer to heap node, index of heap node}
	//   score: score of heap node
	hh := minHeap{heap{
		nodes: []node{
			{
				key:   piTup{pter: &(h.nodes[0]), idx: 0},
				score: h.nodes[0].score,
			},
		},
	}}

	for i := 0; i < n; i++ {
		topIdx = hh.nodes[0].key.(piTup).idx
		k1, ok1 = h.getLeftChild(topIdx)
		k2, ok2 = h.getRightChild(topIdx)

		if ok1 {
			sc1 = h.nodes[k1].score
			// Store pointer to left child of h main node in left child of hh main node
			hh.insertNode(node{key: piTup{pter: &(h.nodes[k1]), idx: k1}, score: sc1})
		}

		if ok2 {
			sc2 = h.nodes[k2].score
			hh.insertNode(node{key: piTup{pter: &(h.nodes[k2]), idx: k2}, score: sc2})
		}

		popped := *(hh.extractTopNode()).key.(piTup).pter
		out = append(out, popped)
	}

	return DumpKeys(out), DumpScores(out)
}

// If only node i violates min-heap criterium, fix it
func (h *minHeap) HeapifyUp(i int) {
	j := h.getParent(i)
	for h.nodes[j].score > h.nodes[i].score {
		h.swapNodes(i, j)
		i, j = j, h.getParent(j)
	}
}

// If top node violates min-heap criterium, fix the heap
func (h *minHeap) HeapifyDown(i int) {
	k, ok := h.getSmallestChild(i)
	var p node
	var c node

	for ok { // Stop if there are no children left
		p = h.nodes[i]
		c = h.nodes[k]

		if c.score < p.score {
			h.swapNodes(i, k)
			i = k
			k, ok = h.getSmallestChild(k)

		} else { // Or if all children are bigger than their parents
			break
		}
	}
}

// Heap Sort
// ---------
// Sort the minHeap and return slice of sorted keys and slice of respective scores
func (h *minHeap) Sort() ([]interface{}, []int) {
	sortedNodes := h.nodes
	for range sortedNodes {
		h.DeleteTop()
	}
	return DumpKeys(sortedNodes), DumpScores(sortedNodes)
}
