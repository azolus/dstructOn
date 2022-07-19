/*

Copyright (C) 2022 Silas Happ <zilas@tutanota.com>
--------------------------------------------------
Use of this source code is governed by the GNU-GPLv3 license that can be found
in the LICENSE file. See also <http://www.gnu.org/licenses/>.

TODO: Migrate from node as a struct to just having a key and a score slice per
heap. Using nodes as tuples was a really impractical idea that leads to a lot
of copying values around.

*/

package heap

import (
	"fmt"
)

type heap struct {
	nodes []node
}

// Tuple of pointer to a node and its index in respectve heap
type piTup struct {
	pter *node
	idx  int
}

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

// Get size of heap
func (h *heap) GetSize() int {
	return len(h.nodes)
}

// Print contents of heap
func (h *heap) Print() {
	fmt.Println("Heap: { (node  score) ... }")
	for _, v := range h.nodes {
		fmt.Printf("%s  %d\n", v.key, v.score)
	}
}
