package queue

import "sync"

type Queue struct {
	lock  sync.Mutex
	items []interface{}
	limit int
}

func newQueue() *Queue {
	return &Queue{
		sync.Mutex{},
		make([]interface{}, 0, 10),
		10,
	}
}

func (q *Queue) Push(item interface{}) {
	length := len(q.items)
	if length > q.limit {
		q.items = q.items[1:length]
		q.items = append(q.items, item)
		return
	}

	q.items = append(q.items)

}

func (q *Queue) Pop() {

}

func (q *Queue) Length() int {
	return len(q.items)
}
