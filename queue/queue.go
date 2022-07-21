package queue

import "log"

/*

Copyright (C) 2022 Silas Happ <zilas@tutanota.com>
--------------------------------------------------
Use of this source code is governed by the GNU-GPLv3.0 license that can be found
in the LICENSE file. See also <http://www.gnu.org/licenses/>.


Purpose
-------
This is a small library implementing a queue as a modified linked list. This is
mainly for fun/learning on my part and will probably be of very little real
use, due to performance issues.

*/

type node struct {
	data interface{}
	next *node
}

type queue struct {
	head   *node
	tail   *node
	length int
}

func NewQueue() *queue {
	return &queue{length: 0}
}

func (q *queue) Length() int {
	return q.length
}

func (q *queue) Enqueue(data interface{}) {
	n := node{data: data}
	if q.length > 0 {
		q.tail.next = &n
	} else {
		q.head = &n
	}
	q.tail = &n
	q.length++
}

func (q *queue) Dequeue() interface{} {
	if q.length == 0 {
		log.Fatal("Can't dequeue from an empty queue...")
	}
	dequeued := q.head
	q.head = q.head.next
	q.length--
	return dequeued
}
