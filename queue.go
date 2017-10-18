package queue

import (
	"github.com/intdxdt/list"
)

//Queue data structure FIFO
type Queue struct {
	list *list.List
}

//NewQueue create a new queue
func NewQueue() *Queue {
	return &Queue{list: list.NewList()}
}

//Size of the queue
func (q *Queue) Size() int {
	return q.list.Len()
}

//Append - add item to the end of queue
func (q *Queue ) Append(item interface{}) *Queue{
	q.list.Append(item)
	return q
}

//AppendLeft - add item to queue from the beginning
func (q *Queue ) AppendLeft(item interface{}) *Queue{
	q.list.AppendLeft(item)
	return q
}

//PopLeft the queue for the first item
func (q *Queue ) PopLeft() interface{} {
	return q.list.PopLeft()
}

//First Looks at the item infront
func (q *Queue) First() interface{} {
	return q.list.First()
}

//Last returns the last item in the queue
func (q *Queue) Last() interface{} {
	return q.list.Last()
}

//Empty - empties the queue
func (q *Queue) Empty() *Queue {
	q.list.Clear()
	return q
}

//IsEmpty check if queue is empty ?
func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}

//Stringer interface
func (q *Queue) String() string {
	return q.list.String()
}

