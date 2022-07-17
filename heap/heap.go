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
each occurence of []Node with []int.

Nodes are key-value pairs, that allow us to effectively sort strings by
assigning a value to them. This can be implemented in whichever way suits your
problem, eg you can score by length, occurences of the character a or any
other helpful sorting criterium.

TODO: Migrate from node as a struct to just having a key and a score slice per
heap. Using nodes as tuples was a really impractical idea that leads to a lot
of copying values around.
*/

package heap

import (
	"log"
)

type node struct {
	key   interface{}
	score int
}

// Tuple of pointer to a node and its index in respectve heap
type piTup struct {
	pter *node
	idx  int
}

type heap struct {
	nodes []node
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

// Create empty node
func NewNode() *node {
	return &node{}
}

// Initialize new node with values
func InitNode(key interface{}, score int) node {
	return node{key: key, score: score}
}

// Get key-field of node
func (n *node) GetKey() interface{} {
	return n.key
}

// Get score-field of node
func (n *node) GetScore() interface{} {
	return n.score
}

// Set key-field of node
func (n *node) SetKey(key interface{}) {
	n.key = key
}

// Set score-field of node
func (n *node) SetScore(score int) {
	n.score = score
}

// Initialize new node slice from interface{} array and score func
func InitNodes(interSl []interface{}, scoreFunc func(interface{}) int) []node {
	nodeSl := make([]node, len(interSl))
	for i, v := range interSl {
		nodeSl[i] = InitNode(v, scoreFunc(v))
	}
	return nodeSl
}

// Dump out all node.keys in a slice
func DumpKeys(nodeSl []node) []interface{} {
	interSl := make([]interface{}, len(nodeSl))
	for i, v := range nodeSl {
		interSl[i] = v.key
	}
	return interSl
}

// Dump out all node.scores in a slice
func DumpScores(nodeSl []node) []int {
	intSl := make([]int, len(nodeSl))
	for i, v := range nodeSl {
		intSl[i] = v.score
	}
	return intSl
}

// ============================================================================
// Heap
// ============================================================================

// For node with index i, return index of parent node
func (*heap) getParent(i int) int {
	return (i - 1) / 2
}

// For node with index i, return index of left child node, if child exists
// Will return -1 and false if there is no child
func (h *heap) getLeftChild(i int) (int, bool) {
	j := 2*i + 1
	if j < len(h.nodes) {
		return j, true
	}
	return -1, false
}

// For node with index i, return index of right child node and true, if child exist
// Will return -1 and false if there is no child
func (h *heap) getRightChild(i int) (int, bool) {
	j := 2*i + 2

	if j < len(h.nodes) {
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
		if h.nodes[k1].score >= h.nodes[k2].score {
			return k2, true
		}
		return k1, true

	} else if ok1 { // Only one child remains
		return k1, true
	} else { // node has no children
		return -1, false
	}
}

// Return the largest child in of a node, if it exists
// Will return -1 and false if there is no child
func (h *heap) getLargestChild(i int) (int, bool) {
	k1, ok1 := h.getLeftChild(i)
	k2, ok2 := h.getRightChild(i)

	if ok2 { // Both children exist
		if h.nodes[k1].score >= h.nodes[k2].score {
			return k1, true
		}
		return k2, true

	} else if ok1 { // Only one child remains
		return k1, true
	} else { // node has no children
		return -1, false
	}
}

// Swap nodes in place i and j
func (h *heap) swapNodes(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// ============================================================================
// Min Heap
// ============================================================================

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
func (h *minHeap) InsertNode(n node) {
	h.nodes = append(h.nodes, n)
	h.HeapifyUp(len(h.nodes) - 1)
}

// Extract top node, fill blank with last node and re-heapify
func (h *minHeap) extractTopNode() node {
	l := len(h.nodes) - 1
	if l < 0 {
		log.Fatal("Can't extract from empty heap...")
	}

	ex := h.nodes[0]
	h.nodes[0] = h.nodes[l]
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
func (h *minHeap) GetTop(n int) ([]interface{}, []int) {
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
			hh.InsertNode(node{key: piTup{pter: &(h.nodes[k1]), idx: k1}, score: sc1})
		}

		if ok2 {
			sc2 = h.nodes[k2].score
			hh.InsertNode(node{key: piTup{pter: &(h.nodes[k2]), idx: k2}, score: sc2})
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

// ============================================================================
// Max Heap
// ============================================================================

// Create an empty maxHeap
func NewMaxHeap(arr []node) *maxHeap {
	h := maxHeap{heap{}}
	return &h
}

// Set nodes field of maxHeap
func (h *maxHeap) SetNodes(interSl []interface{}, scoreFunc func(interface{}) int) {
	h.nodes = InitNodes(interSl, scoreFunc)
	l := len(h.nodes)

	for i := l / 2; i >= 0; i-- {
		h.HeapifyDown(i)
	}
}

// If only node i violates max-heap criterium, fix it
func (h *maxHeap) HeapifyUp(i int) {
	j := h.getParent(i)

	for h.nodes[j].score < h.nodes[i].score {
		h.swapNodes(i, j)
		i, j = j, h.getParent(j)
	}
}

// If top node violates max-heap criterium, fix the heap
func (h *maxHeap) HeapifyDown(i int) {
	k, ok := h.getLargestChild(i)
	var p node
	var c node

	for ok { // Stop if there are no children left
		p = h.nodes[i]
		c = h.nodes[k]

		if c.score > p.score {
			h.swapNodes(i, k)
			i = k
			k, ok = h.getLargestChild(k)

		} else { // Or if all children are smaller than their parents
			break
		}
	}
}

// Extract top node, fill blank with last node and re-heapify
func (h *maxHeap) extractTopNode() node {
	l := len(h.nodes) - 1
	if l < 0 {
		log.Fatal("Can't extract from empty heap...")
	}

	ex := h.nodes[0]
	h.nodes[0] = h.nodes[l]
	h.nodes = h.nodes[:l]
	h.HeapifyDown(0)

	return ex
}

// Extract top node via extractTopNode(...) and return its key and score.
func (h *maxHeap) ExtractTop() (interface{}, int) {
	exNode := h.extractTopNode()
	return exNode.key, exNode.score
}

// Insert node at last position, re-heapify
func (h *maxHeap) InsertNode(n node) {
	h.nodes = append(h.nodes, n)
	h.HeapifyUp(len(h.nodes) - 1)
}

// Return top n elements
func (h *maxHeap) GetTop(n int) []interface{} {
	out := []node{}
	var (
		k1, k2   int  // index of [left:1|right:2] child node
		sc1, sc2 int  // score of [left:1|right:2] child node
		topIdx   int  // index of current top node of hh as indexed in original heap
		ok1, ok2 bool // ok if children exist
	)

	// hh (heap of heaps) stores nodes which contain information about the original heap:
	//   Key: {pointer to heap node, index of that heap node in original heap}
	//   Score: score of heap node in original heap
	hh := maxHeap{heap{
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
			hh.InsertNode(node{key: piTup{pter: &(h.nodes[k1]), idx: k1}, score: sc1})
		}

		if ok2 {
			sc2 = h.nodes[k2].score
			hh.InsertNode(node{key: piTup{pter: &(h.nodes[k2]), idx: k2}, score: sc2})
		}

		popped := *(hh.extractTopNode()).key.(piTup).pter
		out = append(out, popped)
	}

	return DumpKeys(out)
}
