/*
===============================================================================
Copyright (C) 2022 Silas Happ <zilas@tutanota.com>
===============================================================================
Use of this source code is governed by the GNU-GPLv3 license that can be found
in the LICENSE file. See also <http://www.gnu.org/licenses/>.


Purpose
-------
This is a small library implementing the linked-list datastructure for
implementing stacks and the likes.
*/

package linkedList

import (
	"log"
)

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

func (node *node) SetData(data interface{}) {
	node.data = data
}

type linkedList struct {
	head *node
	len  int
}

func NewLinkedList() *linkedList {
	return &linkedList{}
}

func (l *linkedList) GetLength() int {
	return l.len
}

func (l *linkedList) Prepend(data interface{}) {
	second := l.head
	l.head = NewNode()
	l.head.SetData(data)
	l.head.SetNext(second)
	l.len++
}

func (l *linkedList) Peek() interface{} {
	if l.len == 0 {
		log.Fatal("Can't peek into empty linked-list.")
	}
	return l.head.data
}

func (l *linkedList) RemoveFirst() {
	if l.len == 0 {
		log.Fatal("Can't peek into empty linked-list.")
	}
	l.head = l.head.next
	l.len--
}

func (l *linkedList) ExtractFirst() interface{} {
	first := l.head.data
	l.RemoveFirst()
	return first
}
