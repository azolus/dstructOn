/*

Copyright (C) 2022 Silas Happ <zilas@tutanota.com>
--------------------------------------------------
Use of this source code is governed by the GNU-GPLv3.0 license that can be found
in the LICENSE file. See also <http://www.gnu.org/licenses/>.


Purpose
-------
This is a small library implementing the linked-list datastructure for
implementing stacks and the likes.

*/

package linked_list

import (
	"log"
)

// `node` are the elements of a `linkedList`
type node struct {
	data interface{}
	next *node
}

// Create a new node and return its pointer
func NewNode() *node {
	return &node{}
}

// Set the pointer to next node of given node
func (node *node) SetNext(next *node) {
	node.next = next
}

// Set a nodes data
func (node *node) SetData(data interface{}) {
	node.data = data
}

type linkedList struct {
	head *node
	len  int
}

// Create empty `linkedList` and return pointer
func NewLinkedList() *linkedList {
	return &linkedList{}
}

// Return the length of a `linkedList`
func (l *linkedList) GetLength() int {
	return l.len
}

// Prepend (add a node) to `linkedList`
func (l *linkedList) Prepend(data interface{}) {
	second := l.head
	l.head = NewNode()
	l.head.SetData(data)
	l.head.SetNext(second)
	l.len++
}

// Return data of head `node` of `linkedList`
func (l *linkedList) Peek() interface{} {
	if l.len == 0 {
		log.Fatal("Can't peek into empty linked-list.")
	}
	return l.head.data
}

// Delete head `node` of `linkedList`
func (l *linkedList) DeleteHead() {
	if l.len == 0 {
		log.Fatal("Can't peek into empty linked-list.")
	}
	l.head = l.head.next
	l.len--
}

// Return head `node` and remove it from `linkedList`
func (l *linkedList) ExtractHead() interface{} {
	if l.len == 0 {
		log.Fatal("Can't peek into empty linked-list.")
	}
	first := l.head.data
	l.head = l.head.next
	l.len--
	return first
}
