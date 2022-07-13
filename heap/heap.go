/*
===============================================================================
Copyright (C) 2022 Silas Happ <zilas@tutanota.com>
===============================================================================
Use of this source code is governed by the GNU-GPLv3 license that can be found
in the LICENSE file. See also <http://www.gnu.org/licenses/>.


Purpose
-------
This is a small library implementing the heap datastructure for heapsort
and the like. If you just want to create heaps for integers, you can replace
each occurence of `[]Node` with `int[]`.

`Nodes` are key-value pairs, that allow us to effectively sort strings by
assigning a value to them. This can be implemented in whichever way suits your
problem, eg you can score by length, occurences of the character `a` or any
other helpful sorting criterion.
*/

package heap

import (
	"log"
)

type Node struct {
	Key   interface{}
	Score int
	// Info  interface{}
}

type pointIdx struct {
	pointer *Node
	idx     int
}

type heap struct {
	Nodes []Node
}

type minHeap struct {
	heap
}

type maxHeap struct {
	heap
}

// ============================================================================
// Node
// ============================================================================

func NewNode() *Node {
	return &Node{}
}

// ============================================================================
// Min Heap
// ============================================================================

// Create a min heap out of an unsorted `[]Node` array
func NewMinHeap() *minHeap {
	h := minHeap{heap{}}
	return &h
}

func (h *minHeap) SetNodes(arr []Node) {
	h.Nodes = arr
	l := len(h.Nodes)

	for i := l / 2; i >= 0; i-- {
		h.HeapifyDown(i)
	}
}

// Insert node at last position, re-heapify
func (h *minHeap) InsertNode(n Node) {
	h.Nodes = append(h.Nodes, n)
	h.HeapifyUp(len(h.Nodes) - 1)
}

// Extract top node, fill blank with last node and re-heapify
func (h *minHeap) ExtractTop() Node {
	l := len(h.Nodes) - 1
	var ex Node
	if l < 0 {
		log.Fatal("Can't extract from empty heap...")
	}

	ex = h.Nodes[0]
	h.Nodes[0] = h.Nodes[l]
	h.Nodes = h.Nodes[:l]
	h.HeapifyDown(0)

	return ex
}

// Return top n elements
func (h *minHeap) GetTop(n int) []Node {
	out := []Node{}
	var (
		k1, k2, sc1, sc2, topIdx int
		ok1, ok2                 bool
	)

	// heapheap stores nodes which contain information about the original heap:
	//   Key: {pointer to heap node, index of heap node}
	//   Score: score of heap node
	hh := minHeap{heap{
		Nodes: []Node{
			{
				Key:   pointIdx{pointer: &(h.Nodes[0]), idx: 0},
				Score: h.Nodes[0].Score,
			},
		}}}

	for i := 0; i < n; i++ {
		// fmt.Println("pre-insert", hh.Nodes)
		topIdx = hh.Nodes[0].Key.(pointIdx).idx
		k1, ok1 = h.getLeftChild(topIdx)
		k2, ok2 = h.getRightChild(topIdx)

		if ok1 {
			sc1 = h.Nodes[k1].Score
			// Store pointer to left child of h main node in left child of hh main node
			hh.InsertNode(Node{Key: pointIdx{pointer: &(h.Nodes[k1]), idx: k1}, Score: sc1})
		}
		if ok2 {
			sc2 = h.Nodes[k2].Score
			hh.InsertNode(Node{Key: pointIdx{pointer: &(h.Nodes[k2]), idx: k2}, Score: sc2})
		}
		// fmt.Println("post-insert", hh.Nodes)

		popped := *(hh.ExtractTop()).Key.(pointIdx).pointer
		// fmt.Println("Popped", popped)
		// fmt.Println("post-pop", hh.Nodes)
		out = append(out, popped)
	}

	return out
}

// If only node i violates min-heap criterion, fix it
func (h *minHeap) HeapifyUp(i int) {
	j := h.getParent(i)

	for h.Nodes[j].Score > h.Nodes[i].Score {
		h.swapNodes(i, j)
		i, j = j, h.getParent(j)
	}
}

// If top node violates min-heap criterion, fix the heap
func (h *minHeap) HeapifyDown(i int) {
	k, ok := h.getSmallestChild(i)
	var (
		p Node
		c Node
	)

	for ok { // Stop if there are no children left
		p = h.Nodes[i]
		c = h.Nodes[k]
		if c.Score < p.Score {
			h.swapNodes(i, k)
			i = k
			k, ok = h.getSmallestChild(k)

		} else { // Or if all children are bigger than their parents
			break
		}
	}
}

// For node with index i, return index of parent node
func (*heap) getParent(i int) int {
	return (i - 1) / 2
}

// For node with index i, return index of left child node, if child exists
// Will return -1 and false if there is no child
func (h *heap) getLeftChild(i int) (int, bool) {
	j := 2*i + 1

	if j < len(h.Nodes) {
		return j, true
	}

	return -1, false
}

// For node with index i, return index of right child node and true, if child exist
// Will return -1 and false if there is no child
func (h *heap) getRightChild(i int) (int, bool) {
	j := 2*i + 2

	if j < len(h.Nodes) {
		return j, true
	}

	return -1, false
}

// Return the smallest child in of a node, if it exists
// Will return -1 and false if there is no child
func (h *heap) getSmallestChild(i int) (int, bool) {
	k1, ok1 := h.getLeftChild(i)
	k2, ok2 := h.getRightChild(i)

	if ok2 { // Both children exist
		if h.Nodes[k1].Score >= h.Nodes[k2].Score {
			return k2, true
		}
		return k1, true

	} else if ok1 { // Only one child remains
		return k1, true

	} else { // Node has no children
		return -1, false
	}
}

// Return the largest child in of a node, if it exists
// Will return -1 and false if there is no child
func (h *heap) getLargestChild(i int) (int, bool) {
	k1, ok1 := h.getLeftChild(i)
	k2, ok2 := h.getRightChild(i)

	if ok2 { // Both children exist
		if h.Nodes[k1].Score >= h.Nodes[k2].Score {
			return k1, true
		}
		return k2, true

	} else if ok1 { // Only one child remains
		return k1, true

	} else { // Node has no children
		return -1, false
	}
}

func (h *heap) swapNodes(i, j int) {
	h.Nodes[i], h.Nodes[j] = h.Nodes[j], h.Nodes[i]
}

// ============================================================================
// Max Heap
// ============================================================================

// Create a max heap out of an unsorted `[]Node` array
func CreateMaxHeap(arr []Node) maxHeap {
	h := maxHeap{heap{arr}}
	l := len(h.Nodes)

	for i := l / 2; i >= 0; i-- {
		h.HeapifyDown(i)
	}

	return h
}

// Create a min heap out of an unsorted `[]Node` array
func NewMaxHeap(arr []Node) *maxHeap {
	h := maxHeap{heap{}}
	return &h
}

func (h *maxHeap) SetNodes(arr []Node) {
	h.Nodes = arr
	l := len(h.Nodes)

	for i := l / 2; i >= 0; i-- {
		h.HeapifyDown(i)
	}
}

// If only node i violates max-heap criterion, fix it
func (h *maxHeap) HeapifyUp(i int) {
	j := h.getParent(i)

	for h.Nodes[j].Score < h.Nodes[i].Score {
		h.swapNodes(i, j)
		i, j = j, h.getParent(j)
	}
}

// If top node violates max-heap criterion, fix the heap
func (h *maxHeap) HeapifyDown(i int) {
	k, ok := h.getLargestChild(i)
	var (
		p *Node
		c *Node
	)

	for ok { // Stop if there are no children left
		p = &(h.Nodes[i])
		c = &(h.Nodes[k])
		if c.Score > p.Score {
			h.swapNodes(i, k)
			i = k
			k, ok = h.getLargestChild(k)

		} else { // Or if all children are smaller than their parents
			break
		}
	}
}

// Extract top node, fill blank with last node and re-heapify
func (h *maxHeap) ExtractTop() Node {
	l := len(h.Nodes) - 1
	var ex Node
	if l < 0 {
		log.Fatal("Can't extract from empty heap...")
	}

	ex = h.Nodes[0]
	h.Nodes[0] = h.Nodes[l]
	h.Nodes = h.Nodes[:l]
	h.HeapifyDown(0)

	return ex
}

// Insert node at last position, re-heapify
func (h *maxHeap) InsertNode(n Node) {
	h.Nodes = append(h.Nodes, n)
	h.HeapifyUp(len(h.Nodes) - 1)
}

// Return top n elements
func (h *maxHeap) GetTop(n int) []Node {
	out := []Node{}
	var (
		k1, k2, sc1, sc2, topIdx int
		ok1, ok2                 bool
	)

	// heapheap stores nodes which contain information about the original heap:
	//   Key: {pointer to heap node, index of heap node}
	//   Score: score of heap node
	hh := maxHeap{heap{
		Nodes: []Node{
			{
				Key:   pointIdx{pointer: &(h.Nodes[0]), idx: 0},
				Score: h.Nodes[0].Score,
			},
		}}}

	for i := 0; i < n; i++ {
		// fmt.Println("pre-insert", hh.Nodes)
		topIdx = hh.Nodes[0].Key.(pointIdx).idx
		k1, ok1 = h.getLeftChild(topIdx)
		k2, ok2 = h.getRightChild(topIdx)

		if ok1 {
			sc1 = h.Nodes[k1].Score
			// Store pointer to left child of h main node in left child of hh main node
			hh.InsertNode(Node{Key: pointIdx{pointer: &(h.Nodes[k1]), idx: k1}, Score: sc1})
		}
		if ok2 {
			sc2 = h.Nodes[k2].Score
			hh.InsertNode(Node{Key: pointIdx{pointer: &(h.Nodes[k2]), idx: k2}, Score: sc2})
		}
		// fmt.Println("post-insert", hh.Nodes)

		popped := *(hh.ExtractTop()).Key.(pointIdx).pointer
		// fmt.Println("Popped", popped)
		// fmt.Println("post-pop", hh.Nodes)
		out = append(out, popped)
	}

	return out
}
